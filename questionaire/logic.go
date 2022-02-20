package questionaire

import (
	"context"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	db     Database
	logger log.Logger
}

func NewService(db Database, logger log.Logger) Service {
	return &service{
		db:     db,
		logger: logger,
	}
}

func (s service) CreateSet(ctx context.Context, name string, cat string) (string, error) {
	logger := log.With(s.logger, "method", "CreateSet")

	uuid, _ := uuid.NewV4()
	id := uuid.String()

	if cat == "" {
		cat = "general"
	}

	set := Set{
		ID:       id,
		Name:     name,
		Category: cat,
	}

	if err := s.db.CreateSet(ctx, set); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create set", id)

	return "Success", nil
}

func (s service) GetSet(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetSet")

	name, err := s.db.GetSet(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get set", id)

	return name, nil
}
