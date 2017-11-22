package handler

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/pkg/errors"
)

func parseRequestParam(r *http.Request, request interface{}) error {
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

	// 解析url命令行参数
	urlValues, err := ParseURL(r)
	if err != nil {
		return errors.Wrapf(BeanParamError, "parse url error, %s", err.Error())
	}
	// 遍历每个域
	for _, field := range parseRequestFields(typ) {
		if field.loc != "param" {
			continue
		}
		// 取出来要设置的值
		input := urlValues.GetString(field.fieldName)
		if input == "" && field.required {
			return errors.Wrapf(BeanParamError, "field [%s] must be set", field.fieldName)
		}
		// 取出来结构中要设置的域到v
		if err := input2Value(input, v.FieldByName(field.name)); err != nil {
			return errors.Wrap(BeanParamError, fmt.Sprintf("field [%s] %s", field.fieldName, err.Error()))
		}
	}
	return nil
}
