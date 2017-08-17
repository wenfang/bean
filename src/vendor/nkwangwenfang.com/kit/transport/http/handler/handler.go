package handler

import (
	"context"
	"net/http"
	"reflect"
	"runtime/debug"

	"nkwangwenfang.com/kit/endpoint"
	"nkwangwenfang.com/log"
)

type handler struct {
	ctx context.Context
	e   endpoint.Endpoint
	typ reflect.Type
}

type handler1 struct {
	e   endpoint.Endpoint
	typ reflect.Type
}

// New 创建一个http.Handler, request为请求的结构，类型必须为struct或者slice
func New(ctx context.Context, e endpoint.Endpoint, request interface{}) http.Handler {
	return &handler{ctx: ctx, e: e, typ: reflect.TypeOf(request)}
}

func New1(e endpoint.Endpoint, request interface{}) http.Handler {
	// 检查request的类型必须为struct,否则无法启动
	typ := reflect.TypeOf(request)
	if typ.Kind() != reflect.Struct {
		panic("request must be struct")
	}
	return &handler1{e: e, typ: typ}
}

func (h *handler1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理请求内部调用panic的情况
	defer func() {
		if errData := recover(); errData != nil {
			log.Error(errData, "stack", string(debug.Stack()))
			Output1(w, BeanInnerError, nil)
		}
	}()

	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 创建请求，类型为interface{}，内部必须为一个指针,指向struct
	//request := reflect.New(h.typ).Interface()
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理panic情况
	defer func() {
		if errData := recover(); errData != nil {
			log.Error(errData, "stack", string(debug.Stack()))
			OutputError(w, errData, ErrorTypeInner)
		}
	}()

	// 类型必须为struct或slice
	if h.typ.Kind() != reflect.Struct && h.typ.Kind() != reflect.Slice {
		log.Error("request must be struct or slice")
		OutputError(w, "request must be struct or slice", ErrorTypeInner)
		return
	}
	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 创建请求，类型为interface{}，内部必须为一个指针,指向slice或struct
	request := reflect.New(h.typ).Interface()
	// 解析HTTP请求，来自请求路径
	if errData, errType := parseRequestPath(r, request); errType != nil {
		OutputError(w, errData, errType)
		return
	}
	// 解析HTTP请求,来自url参数
	if errData, errType := parseRequestParam(r, request); errType != nil {
		OutputError(w, errData, errType)
		return
	}
	// 解析HTTP请求,来自JSON请求体
	if errData, errType := parseRequestBody(r, request); errType != nil {
		OutputError(w, errData, errType)
		return
	}
	// 执行请求处理，生成响应
	response, errType := h.e(h.ctx, request)
	if errType != nil {
		OutputError(w, response, errType)
		return
	}
	// 输出响应
	if errType := OutputResponse(w, response); errType != nil {
		OutputError(w, "encode response error", errType)
	}
}

// NotFound 未发现对应的请求实体
func NotFound(w http.ResponseWriter, r *http.Request) {
	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	OutputError(w, nil, ErrorTypeEntity)
}

func NotFound1(w http.ResponseWriter, r *http.Request) {
	// 设置缺省返回的Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	Output1(w, BeanEntityError, nil)
}
