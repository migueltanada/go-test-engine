package main

import (
	"context"
	"flag"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"go-questionaire/questionaire"

	"go.mongodb.org/mongo-driver/mongo"
)

// const dbsource = "postgresql://postgres:example@localhost:5432/questionaire?sslmode=disable"
const dbSource = "mongodb://root:example@localhost:27017"

func main() {
	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *mongo.Database
	{
		db = GetMongoDB(context.TODO(), dbSource, "questionaire")
	}
	flag.Parse()
	ctx := context.Background()
	var srv questionaire.Service
	{
		repository := questionaire.NewDb(db, logger)

		srv = questionaire.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := questionaire.MakeEndpoints(srv)

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := questionaire.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}
