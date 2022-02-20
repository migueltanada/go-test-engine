package questionaire

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/set").Handler(httptransport.NewServer(
		endpoints.CreateSet,
		decodeSetReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/set/{setId}").Handler(httptransport.NewServer(
		endpoints.GetSet,
		decodeSetNameReq,
		encodeResponse,
	))

	r.Methods("POST").Path("/question").Handler(httptransport.NewServer(
		endpoints.CreateQna,
		decodeQnaReq,
		encodeResponse,
	))

	r.Methods("GET").Path("/question/{qnaId}").Handler(httptransport.NewServer(
		endpoints.GetQna,
		decodeQnaQuestionReq,
		encodeResponse,
	))

	return r

}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
