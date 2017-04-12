package handler

import (
	"reflect"
)

type requestFields struct {
	loc       string
	name      string
	fieldName string
	typ       reflect.Type
	required  bool
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
