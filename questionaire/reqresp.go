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
		Id string `json:"id"`
	}
	GetSetResponse struct {
		Name string `json:"name"`
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

func decodeNameReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req GetSetRequest
	vars := mux.Vars(r)

	req = GetSetRequest{
		Id: vars["id"],
	}
	return req, nil
}
