package iterator

import (
	"iter"

	"github.com/lord-server/panorama/internal/util/result"
)

type ResultSeq[T any] iter.Seq[result.Result[T]]
