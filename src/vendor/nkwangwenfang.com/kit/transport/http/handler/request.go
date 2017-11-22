package handler

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// input2ValueBasic 将输入的input的值写入value中,只支持基本类型
func input2ValueBasic(input string, v reflect.Value) error {
	// 获取值的类型
	kind := v.Kind()
	// 根据kind的类型进行设置
	switch kind {
	case reflect.Bool:
		vBool, err := strconv.ParseBool(input)
		if err != nil {
			return fmt.Errorf("parse bool error")
		}
		v.SetBool(vBool)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		vInt, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			return fmt.Errorf("parse int error")
		}
		v.SetInt(vInt)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		vUint, err := strconv.ParseUint(input, 10, 64)
		if err != nil {
			return fmt.Errorf("parse uint error")
		}
		v.SetUint(vUint)
	case reflect.Float32, reflect.Float64:
		vFloat, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return fmt.Errorf("parse float error")
		}
		v.SetFloat(vFloat)
	case reflect.String:
		v.SetString(input)
	default:
		return fmt.Errorf("type [%s] not support", kind)
	}
	return nil
}

// input2Value 将输入的input值写入vaule中
func input2Value(input string, v reflect.Value) error {
	// 如果值不能被设置,返回错误
	if !v.CanSet() {
		return fmt.Errorf("field can not be set")
	}
	// 根据值的类型进行处理
	switch v.Kind() {
	case reflect.Ptr: // 指针类型,创建新元素
		nv := reflect.New(v.Type().Elem())
		v.Set(nv)
		v = nv.Elem()
	case reflect.Slice:
		// split出来每项,分别进行转换
		nv := reflect.MakeSlice(v.Type(), 0, 0)
		for _, item := range strings.Split(input, ",") {
			itemv := reflect.Indirect(reflect.New(v.Type().Elem()))
			if err := input2ValueBasic(item, itemv); err != nil {
				return err
			}
			nv = reflect.Append(nv, itemv)
		}
		v.Set(nv)
		return nil
	}
	return input2ValueBasic(input, v)
}
