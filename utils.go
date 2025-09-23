package go_context

import (
	"context"
)

// CancelContextOnError cancels the given context if an error is encountered.
//
// Parameters:
//
// ctx: The context to be stopped.
// cancelFn: A function that cancels the context with a given error.
// fn: A function that returns an error.
// loggerProducer: The logger producer to log messages.
//
// Returns:
//
// A function that executes the provided function and cancels the context if an error occurs.
func CancelContextOnError(
	ctx context.Context,
	cancelFn context.CancelFunc,
	fn func(context.Context) error,
) func() error {
	return func() error {
		if err := fn(ctx); err != nil {
			if cancelFn != nil {
				cancelFn()
			}
			return err
		}
		return nil
	}
}
