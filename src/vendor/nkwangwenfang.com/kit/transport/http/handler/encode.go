package handler

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"

	"nkwangwenfang.com/log"
)

type response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func EncodeResponse(w http.ResponseWriter, data interface{}) error {
	if err := json.NewEncoder(w).Encode(response{Code: 0, Message: "Success", Data: data}); err != nil {
		log.Error("response data encode error", "error", err)
		return ErrInner
	}
	return nil
}

func EncodeError(w http.ResponseWriter, data interface{}, err error) {
	e, ok := errors.Cause(err).(*Error)
	if !ok {
		e = ErrInner
		log.Error("err invalid, set to ErrInner", "error", err)
	}
	if err := json.NewEncoder(w).Encode(response{Code: e.Code, Message: e.Message, Errors: data}); err != nil {
		log.Error("error data invalid, set to nil", "data", data)
		json.NewEncoder(w).Encode(response{Code: e.Code, Message: e.Message})
	}
}
