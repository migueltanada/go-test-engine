package questionaire

import "context"

type Service interface {
	CreateSet(ctx context.Context, name string, cat string) (string, error)
	GetSet(ctx context.Context, setId string) (string, error)
	CreateQna(ctx context.Context, setId string, question string, choices Choice) (string, error)
	GetQna(ctx context.Context, qnaId string) (string, error)
}
