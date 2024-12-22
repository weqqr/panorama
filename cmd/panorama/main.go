package main

import (
	"os"

	"github.com/alexflint/go-arg"
	"github.com/lord-server/panorama/internal/api/server"
	"github.com/lord-server/panorama/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type FullRenderArgs struct{}

type RunArgs struct{}

var args struct {
	ConfigPath string   `arg:"-c,--config" default:"config.toml"`
	Run        *RunArgs `arg:"subcommand:run"`
}

func main() {
	arg.MustParse(&args)

	loggerConfig := zap.NewDevelopmentConfig()
	loggerConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := loggerConfig.Build()

	config, err := config.Load(args.ConfigPath)
	if err != nil {
		logger.Error("unable to load config",
			zap.String("path", args.ConfigPath),
			zap.Error(err))

		os.Exit(1)
	}

	switch {
	case args.Run != nil:
		err = run(logger, config)

	default:
		logger.Warn("command not specified, proceeding with run")

		err = run(logger, config)
	}

	if err != nil {
		os.Exit(1)
	}
}

func run(logger *zap.Logger, config config.Config) error {
	quit := make(chan bool)

	logger.Info("starting web server",
		zap.String("address", config.Web.ListenAddress))

	go func() {
		server.Serve(logger, &config)
		quit <- true
	}()

	<-quit

	return nil
}
