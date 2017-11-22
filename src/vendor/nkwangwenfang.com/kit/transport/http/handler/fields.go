package handler

import (
	"reflect"
)

// requestFields 解析请求的各个域
type requestFields struct {
	loc       string       // 域所在的位置path上或者param上
	name      string       // 在结构中对应的名称
	fieldName string       // 在path或param上的名称
	typ       reflect.Type // 想转化为的类型
	required  bool         // 是否是必须字段
}

func parseRequestFields(t reflect.Type) []requestFields {
	var fields []requestFields
	for i := 0; i < t.NumField(); i++ {
		// 解析req_path
		tag := t.Field(i).Tag.Get("req_path")
		if tag != "" {
			fieldName, _ := parseTag(tag)
			fields = append(fields, requestFields{
				loc:       "path",
				name:      t.Field(i).Name,
				fieldName: fieldName,
				typ:       t.Field(i).Type,
			})
		}
		// 解析req_param
		tag = t.Field(i).Tag.Get("req_param")
		if tag != "" {
			fieldName, opt := parseTag(tag)
			fields = append(fields, requestFields{
				loc:       "param",
				name:      t.Field(i).Name,
				fieldName: fieldName,
				typ:       t.Field(i).Type,
				required:  opt.Contains("required"),
			})
		}
	}
	return fields
}
