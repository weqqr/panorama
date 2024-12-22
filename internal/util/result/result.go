package result

type Result[T any] struct {
	Value T
	Error error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		Value: value,
	}
}

func Error[T any](err error) Result[T] {
	return Result[T]{
		Error: err,
	}
}
