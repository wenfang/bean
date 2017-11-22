package handler

import (
	"net/http"
	"reflect"
	"runtime/debug"

	"nkwangwenfang.com/kit/endpoint"
	"nkwangwenfang.com/log"
)

type handler struct {
	e   endpoint.Endpoint
	typ reflect.Type
}

// New 创建一个http.Handler, request为请求的结构，类型必须为struct或者slice
func New(e endpoint.Endpoint, request interface{}) http.Handler {
	// 检查request的类型必须为struct,否则无法启动
	typ := reflect.TypeOf(request)
	if typ.Kind() != reflect.Struct {
		panic("request must be struct")
	}
	return &handler{e: e, typ: typ}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 处理请求内部调用panic的情况
	defer func() {
		if errData := recover(); errData != nil {
			log.Error(errData, "stack", string(debug.Stack()))
			Output(w, nil, BeanInnerError)
		}
	}()

	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 创建请求，类型为interface{}，内部必须为一个指针,指向struct
	request := reflect.New(h.typ).Interface()

	// 解析HTTP请求，来自请求路径
	if err := parseRequestPath(r, request); err != nil {
		log.Error(err)
		Output(w, nil, err)
		return
	}
	// 解析HTTP请求,来自url参数
	if err := parseRequestParam(r, request); err != nil {
		log.Error(err)
		Output(w, nil, err)
		return
	}
	// 解析HTTP请求,来自JSON请求体
	if err := parseRequestBody(r, request); err != nil {
		log.Error(err)
		Output(w, nil, err)
		return
	}
	// 执行请求处理，生成响应
	data, err := h.e(r.Context(), request)
	if err != nil {
		log.Error(err)
	}
	Output(w, data, err)
}

// NotFound 未发现对应的请求实体
func NotFound(w http.ResponseWriter, r *http.Request) {
	// 设置缺省的返回Content-Type
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	Output(w, nil, BeanEntityError)
}
