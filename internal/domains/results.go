package domains

// Result represents an operation outcome with success/failure tracking
type Result[T any] struct {
	value   T
	err     *Error
	success bool
}

// Success creates a successful result
func Success[T any](value T) Result[T] {
	return Result[T]{
		value:   value,
		success: true,
	}
}

// Failure creates a failed result with domain error
func Failure[T any](err Error) Result[T] {
	return Result[T]{
		err:     &err,
		success: false,
	}
}

// IsSuccess checks if operation succeeded
func (r Result[T]) IsSuccess() bool {
	return r.success
}

// IsFailure checks if operation failed
func (r Result[T]) IsFailure() bool {
	return !r.success
}

// Value returns the success value (panics if called on failure)
func (r Result[T]) Value() T {
	if !r.success {
		panic("Cannot access value of a failed result")
	}
	return r.value
}

// Error returns the error (nil if success)
func (r Result[T]) Error() *Error {
	return r.err
}

// Unwrap returns value and error (Go-style)
func (r Result[T]) Unwrap() (T, *Error) {
	return r.value, r.err
}

// Map transforms the success value
func (r Result[T]) Map(fn func(T) any) Result[any] {
	if r.IsFailure() {
		return Failure[any](*r.err)
	}
	return Success[any](fn(r.value))
}
