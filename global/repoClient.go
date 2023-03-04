package global

import "context"

type RepoClient interface {
	NewTxContext(ctx context.Context) (context.Context, error)
	CommitTx(ctx context.Context) error
	RollbackTx(ctx context.Context)
}
