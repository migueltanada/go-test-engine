package questionaire

import "context"

type Set struct {
	SetID    string `json:"setId,omitempty"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Database interface {
	CreateSet(ctx context.Context, set Set) error
	GetSet(ctx context.Context, setId string) (string, error)
	CreateQna(ctx context.Context, qna Qna) error
	GetQna(ctx context.Context, qnaId string) (string, error)
}
