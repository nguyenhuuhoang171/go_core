package safego

import (
	"context"
	"log"
	"runtime/debug"
)

// SafeGo runs fn in a new goroutine and recovers from any panic inside it,
// so a bug in fn cannot crash the whole process. name identifies the
// goroutine in logs, and ctx is passed through to fn. The panic value and
// stack trace are logged when a recover happens.
func SafeGo(ctx context.Context, name string, fn func(ctx context.Context)) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[safego] %s: recovered panic: %v\n%s", name, r, debug.Stack())
			}
		}()
		fn(ctx)
	}()
}
