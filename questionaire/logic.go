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
		SetID:    id,
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

func (s service) CreateQna(ctx context.Context, setId string, question string, choices Choice) (string, error) {
	logger := log.With(s.logger, "method", "CreateSet")

	uuid, _ := uuid.NewV4()
	questionId := uuid.String()

	if setId == "" {
		setId = "noset"
	}

	qna := Qna{
		SetID:    setId,
		QnaID:    questionId,
		Question: question,
		Choice:   choices,
	}

	if err := s.db.CreateQna(ctx, qna); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create set", questionId)

	return "Success", nil
}

func (s service) GetQna(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetSet")

	name, err := s.db.GetQna(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get set", id)

	return name, nil
}
