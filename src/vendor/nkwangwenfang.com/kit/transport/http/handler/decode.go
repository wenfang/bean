package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
)

// parseObject 解析对象，必须是一个指向struct的指针结构
func parseObject(v interface{}) (reflect.Value, interface{}) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return reflect.Value{}, &ErrorReason{Reason: "obj is not pointer or is nil"}
	}

	for {
		rv = rv.Elem()
		if rv.Kind() == reflect.Ptr {
			if rv.IsNil() {
				return reflect.Value{}, &ErrorReason{Reason: "obj is nil pointer"}
			}
			continue
		}

		if rv.Kind() != reflect.Struct {
			return reflect.Value{}, &ErrorReason{Reason: "obj is not struct"}
		}
		break
	}

	return rv, nil
}

type field struct {
	name      string
	fieldName string

	typ      reflect.Type
	required bool
}

func parseFields(t reflect.Type) []field {
	var fields []field
	for i := 0; i < t.NumField(); i++ {
		tag := t.Field(i).Tag.Get("httpreq")
		if tag == "" {
			continue
		}
		fieldName, opt := parseTag(tag)
		fields = append(fields, field{
			name:      t.Field(i).Name,
			fieldName: fieldName,
			typ:       t.Field(i).Type,
			required:  opt.Contains("required"),
		})
	}

	return fields
}

func setValue(rv reflect.Value, kind reflect.Kind, fieldName string, fieldValue string) interface{} {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		if fieldValue != "" {
			vInt, err := strconv.ParseInt(fieldValue, 10, 64)
			if err != nil {
				return &ErrorReason{Reason: fmt.Sprintf("field %s must be int", fieldName)}
			}
			rv.SetInt(vInt)
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64:
		if fieldValue != "" {
			vUint, err := strconv.ParseUint(fieldValue, 10, 64)
			if err != nil {
				return &ErrorReason{Reason: fmt.Sprintf("field %s must be uint", fieldName)}
			}
			rv.SetUint(vUint)
		}
	case reflect.Float32, reflect.Float64:
		if fieldValue != "" {
			vFloat, err := strconv.ParseFloat(fieldValue, 64)
			if err != nil {
				return &ErrorReason{Reason: fmt.Sprintf("field %s must be float", fieldName)}
			}
			rv.SetFloat(vFloat)
		}
	case reflect.String:
		rv.SetString(fieldValue)
	default:
		return &ErrorReason{Reason: fmt.Sprintf("field %s type %s not support", fieldName, kind)}
	}
	return nil
}

func decodeRequest(req *http.Request, v interface{}) (interface{}, error) {
	rv, errReason := parseObject(v)
	if errReason != nil {
		return errReason, ErrInner
	}

	for _, field := range parseFields(rv.Type()) {
		fieldValue := req.FormValue(field.fieldName)
		if fieldValue == "" && field.required {
			return &ErrorReason{Reason: fmt.Sprintf("field %s must be set", field.fieldName)}, ErrParameter
		}

		f := rv.FieldByName(field.name)
		if !f.CanSet() {
			continue
		}

		if field.typ.Kind() == reflect.Ptr {
			if fieldValue != "" {
				n := reflect.New(field.typ.Elem())
				if errReason := setValue(n.Elem(), field.typ.Elem().Kind(), field.fieldName, fieldValue); errReason != nil {
					return errReason, ErrParameter
				}
				f.Set(n)
			}
			continue
		}
		if errReason := setValue(f, field.typ.Kind(), field.fieldName, fieldValue); errReason != nil {
			return errReason, ErrParameter
		}
	}
	return nil, nil
}

func decodeRequestBody(req *http.Request, v interface{}) (interface{}, error) {
	//TODO: 解析JSON请求体
	return nil, nil
}
