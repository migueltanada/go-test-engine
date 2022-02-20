package questionaire

import (
	"context"
	"errors"

	"github.com/go-kit/log"
	"go.mongodb.org/mongo-driver/mongo"
)

var ErrDb = errors.New("unable to handle repo request")

type database struct {
	db     *mongo.Database
	logger log.Logger
}

func NewDb(db *mongo.Database, logger log.Logger) Database {
	return &database{
		db:     db,
		logger: log.With(logger, "repo", "mongo"),
	}
}

func (db *database) CreateSet(ctx context.Context, set Set) error {

	collection := db.db.Collection("set")

	if set.Name == "" || set.Category == "" {
		return ErrDb
	}

	_, err := collection.InsertOne(context.TODO(), set)
	if err != nil {
		return ErrDb
	}

	return nil
}

func (db *database) GetSet(ctx context.Context, id string) (string, error) {
	var name string
	// err := db.db.QueryRow("SELECT name FROM set WHERE id=$1", id).Scan(&name)
	// if err != nil {
	// 	return "", DbErr
	// }

	return name, nil
}

func (db *database) CreateQna(ctx context.Context, qna Qna) error {

	collection := db.db.Collection("qna")

	// level.Info(db.logger).Log("err", qna.QnaID)
	// level.Info(db.logger).Log("err", qna.SetID)
	// level.Info(db.logger).Log("err", qna.Question)
	if qna.QnaID == "" || qna.SetID == "" || qna.Question == "" {
		return ErrDb
	}

	_, err := collection.InsertOne(context.TODO(), qna)
	if err != nil {
		return err
	}

	return nil
}

func (db *database) GetQna(ctx context.Context, id string) (string, error) {
	var name string
	// err := db.db.QueryRow("SELECT name FROM set WHERE id=$1", id).Scan(&name)
	// if err != nil {
	// 	return "", DbErr
	// }

	return name, nil
}
