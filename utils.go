package go_context

import (
	"context"
	"fmt"
	"runtime/debug"
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
	return func() (err error) {
        // Recover from panic and return as error
        defer func() {
            if r := recover(); r != nil {
                if cancelFn != nil {
                    cancelFn()
                }
                switch v := r.(type) {
                case error:
                    err = fmt.Errorf("panic recovered: %w\n%s", v, debug.Stack())
                default:
                    err = fmt.Errorf("panic recovered: %v\n%s", v, debug.Stack())
                }
            }
        }()

        if e := fn(ctx); e != nil {
            if cancelFn != nil {
                cancelFn()
            }
            return e
        }
        return nil
	}
}
