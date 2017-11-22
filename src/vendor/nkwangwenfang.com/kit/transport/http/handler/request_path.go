package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func parseRequestPath(r *http.Request, request interface{}) error {
	// 转换成对应的值类型
	v := reflect.ValueOf(request)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.Wrap(BeanInnerError, "request is not pointer or is nil")
	}
	// 获取内层元素
	v = v.Elem()

	typ := v.Type()
	// 需要解析的对象必须是struct类型,如果不是直接返回
	if typ.Kind() != reflect.Struct {
		return nil
	}

	vars := mux.Vars(r)
	for _, field := range parseRequestFields(typ) {
		// 必须为解析路径上的参数
		if field.loc != "path" {
			continue
		}
		// 取出来要设置的值
		input, ok := vars[field.fieldName]
		if !ok {
			continue
		}
		// 取出来结构中要设置的域到v
		if err := input2Value(input, v.FieldByName(field.name)); err != nil {
			return errors.Wrap(BeanPathError, fmt.Sprintf("field [%s] %s", field.fieldName, err.Error()))
		}
	}
	return nil
}
