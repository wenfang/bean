package handler1

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

func parseRequestBody(r *http.Request, request interface{}) error {
	// 只处理POST/PUT/DELETE方法
	if r.Method != "POST" && r.Method != "PUT" && r.Method != "DELETE" {
		return nil
	}
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return errors.Wrap(BeanBodyError, err.Error())
	}
	return nil
}
