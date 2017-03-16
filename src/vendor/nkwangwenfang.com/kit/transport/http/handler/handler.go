package handler

import (
	"context"
	"net/http"
	"reflect"

	"nkwangwenfang.com/kit/endpoint"
)

type handler struct {
	ctx context.Context
	e   endpoint.Endpoint
	t   reflect.Type
}

func New(ctx context.Context, e endpoint.Endpoint, t reflect.Type) http.Handler {
	return &handler{ctx: ctx, e: e, t: t}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.t.Kind() != reflect.Struct {
		EncodeError(w, &ErrorReason{Reason: "obj must be struct"}, ErrInner)
		return
	}
	// 创建请求，类型为interface{}，内部必须为一个指向结构体的指针
	req := reflect.New(h.t).Interface()
	// 解析HTTP请求
	if errReason, err := decodeRequest(r, req); err != nil {
		EncodeError(w, errReason, err)
		return
	}
	// 解析HTTP JSON请求体
	if errReason, err := decodeRequestBody(r, req); err != nil {
		EncodeError(w, errReason, err)
		return
	}
	// 执行请求处理，生成响应
	rsp, err := h.e(h.ctx, req)
	if err != nil {
		EncodeError(w, rsp, err)
		return
	}
	// 输出响应
	if err := EncodeResponse(w, rsp); err != nil {
		EncodeError(w, &ErrorReason{Reason: "encode response error"}, err)
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	EncodeError(w, nil, ErrEntity)
}
