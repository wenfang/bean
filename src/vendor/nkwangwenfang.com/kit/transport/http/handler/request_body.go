package handler

import (
	"encoding/json"
	"net/http"
)

func parseRequestBody(req *http.Request, v interface{}) (interface{}, error) {
	if err := json.NewDecoder(req.Body).Decode(v); err != nil {
		return err.Error(), ErrorBody
	}
	return nil, nil
}
