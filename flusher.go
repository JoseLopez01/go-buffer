package buffer

type (
	// FlusherFunc represents a flush function.
	FlusherFunc[T any] func(items []T)
)

func (fn FlusherFunc[T]) Write(items []T) {
	fn(items)
}
