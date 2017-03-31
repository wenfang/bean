package handler

import (
	"encoding/json"
	"net/http"
)

func parseRequestBody(req *http.Request, v interface{}) (interface{}, error) {
	// 只处理POST/PUT/DELETE方法
	if req.Method != "POST" && req.Method != "PUT" && req.Method != "DELETE" {
		return nil, nil
	}
	if err := json.NewDecoder(req.Body).Decode(v); err != nil {
		return err.Error(), ErrorBody
	}
	return nil, nil
}
