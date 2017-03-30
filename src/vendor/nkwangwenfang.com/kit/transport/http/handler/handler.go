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
	typ reflect.Type
}

// New 创建一个http.Handler, t的类型支持两种，Struct和Slice
func New(ctx context.Context, e endpoint.Endpoint, request interface{}) http.Handler {
	return &handler{ctx: ctx, e: e, typ: reflect.TypeOf(request)}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 类型必须为struct或slice
	if h.typ.Kind() != reflect.Struct && h.typ.Kind() != reflect.Slice {
		OutputError(w, &ErrorReason{Reason: "request must be struct or slice"}, ErrorInner)
		return
	}
	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 创建请求，类型为interface{}，内部必须为一个指针,指向slice或struct
	request := reflect.New(h.typ).Interface()
	if errReason, err := parseRequestPath(r, request); err != nil {
		OutputError(w, errReason, err)
		return
	}
	// 解析HTTP请求,来自url参数
	if errReason, err := parseRequestParam(r, request); err != nil {
		OutputError(w, errReason, err)
		return
	}
	// 解析HTTP请求,来自JSON请求体
	if errReason, err := parseRequestBody(r, request); err != nil {
		OutputError(w, errReason, err)
		return
	}
	// 执行请求处理，生成响应
	response, err := h.e(h.ctx, request)
	if err != nil {
		OutputError(w, response, err)
		return
	}
	// 输出响应
	if err := OutputResponse(w, response); err != nil {
		OutputError(w, &ErrorReason{Reason: "encode response error"}, err)
	}
}

// NotFound 未发现对应的请求实体
func NotFound(w http.ResponseWriter, r *http.Request) {
	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	OutputError(w, nil, ErrorEntity)
}
