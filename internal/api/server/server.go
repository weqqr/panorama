package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lord-server/panorama/internal/config"
	"github.com/lord-server/panorama/static"
	"go.uber.org/zap"
)

const maxRequestSize = 1 * 1024 * 1024 // 1 MiB

func Serve(logger *zap.Logger, config *config.Config) {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.FS(static.UI())))
	mux.Handle("GET /tiles/", http.StripPrefix("/tiles", http.FileServer(http.Dir(config.System.TilesPath))))

	httpServer := &http.Server{
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		Addr:              config.Web.ListenAddress,
		Handler:           mux,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		logger.Error("failed to start web server", zap.Error(err))
	}
}

func handle[Request, Response any](handleFunc func(context.Context, Request) (Response, error)) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var request Request

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		err := decoder.Decode(&request)
		if err != nil {
			renderResponse(w, http.StatusBadRequest, ErrorResponse{
				Message: fmt.Sprintf("malformed json: %v", err),
			})
		}

		defer r.Body.Close()

		response, err := handleFunc(r.Context(), request)
		if err != nil {
			// FIXME:
			renderResponse(w, http.StatusBadRequest, ErrorResponse{
				Message: err.Error(),
			})
		}

		renderResponse(w, http.StatusOK, response)
	}

	return http.MaxBytesHandler(http.HandlerFunc(handler), maxRequestSize)
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func renderResponse(w http.ResponseWriter, code int, response any) {
	data, err := json.Marshal(&response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error encoding response: %v", err)))

		return
	}

	w.WriteHeader(code)
	w.Write(data)
}
