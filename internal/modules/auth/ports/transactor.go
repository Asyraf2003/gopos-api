package ports

import "context"

type Transactor interface {
	RunInTx(ctx context.Context, fn func(context.Context) error) error
}
