package questionaire

import "context"

type Service interface {
	CreateSet(ctx context.Context, name string, cat string) (string, error)
	GetSet(ctx context.Context, id string) (string, error)
}
