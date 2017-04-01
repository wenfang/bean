package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

func parseRequestParam(r *http.Request, v interface{}) (interface{}, error) {
	// 转换成对应的值类型
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Sprintf("request is not pointer or is nil"), ErrorInner
	}
	// 获取内层元素
	rv = rv.Elem()

	typ := rv.Type()
	// 需要解析的对象必须是struct类型,如果不是直接返回
	if typ.Kind() != reflect.Struct {
		return nil, nil
	}

	for _, field := range parseRequestFields(typ) {
		if field.loc != "param" {
			continue
		}
		// 取出来要设置的值
		value := r.FormValue(field.fieldName)
		if value == "" && field.required {
			return fmt.Sprintf("field [%s] must be set", field.fieldName), ErrorParameter
		}
		// 取出来要设置的域
		f := rv.FieldByName(field.name)
		if !f.CanSet() {
			continue
		}

		// 如果是指针类型，创建一个
		kind := field.typ.Kind()
		if kind == reflect.Ptr {
			newf := reflect.New(field.typ.Elem())
			f.Set(newf)
			f = newf.Elem()
			kind = field.typ.Elem().Kind()
		}
		// 根据对应域的类型进行设置
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			vInt, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return fmt.Sprintf("request field [%s] content [%s] parse int error", field.name, value), ErrorParameter
			}
			f.SetInt(vInt)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			vUint, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				return fmt.Sprintf("request field [%s] content [%s] parse uint error", field.name, value), ErrorParameter
			}
			f.SetUint(vUint)
		case reflect.Float32, reflect.Float64:
			vFloat, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return fmt.Sprintf("request field [%s] content [%s] parse float error", field.name, value), ErrorParameter
			}
			f.SetFloat(vFloat)
		case reflect.String:
			f.SetString(value)
		default:
			return fmt.Sprintf("request field [%s] type [%s] not support", field.name, field.typ.Kind()), ErrorParameter
		}
	}
	return nil, nil
}
