package questionaire

import "context"

type Set struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

type Database interface {
	CreateSet(ctx context.Context, set Set) error
	GetSet(ctx context.Context, id string) (string, error)
}
