package questionaire

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreateSetRequest struct {
		Name     string `json:"name"`
		Category string `json:"category"`
	}
	CreateSetResponse struct {
		Ok string `json:"ok"`
	}

	GetSetRequest struct {
		SetID string `json:"setId"`
	}
	GetSetResponse struct {
		Name string `json:"name"`
	}

	CreateQnaRequest struct {
		SetID    string `json:"setId"`
		QnaID    string `json:"questionId,omitempty"`
		Question string `json:"question"`
		Choice   Choice `json:"choices"`
	}
	CreateQnaResponse struct {
		Ok string `json:"ok"`
	}

	GetQnaRequest struct {
		QnaID string `json:"questionId"`
	}
	GetQnaResponse struct {
		Question string `json:"question"`
	}
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeSetReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateSetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeSetNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetSetRequest
	vars := mux.Vars(r)

	req = GetSetRequest{
		SetID: vars["setId"],
	}
	return req, nil
}

func decodeQnaReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req CreateQnaRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func decodeQnaQuestionReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetQnaRequest
	vars := mux.Vars(r)

	req = GetQnaRequest{
		QnaID: vars["questionId"],
	}
	return req, nil
}
