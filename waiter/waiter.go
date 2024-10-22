package waiter

import "context"

type Waiter interface {
	Wait(context.Context) error
}
