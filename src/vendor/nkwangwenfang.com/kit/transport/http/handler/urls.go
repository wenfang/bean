package handler

import (
	"net/http"
	"net/url"
)

// URLValues 代表Url参数值
type URLValues url.Values

// Get 获取名称为key的第一个值
func (u URLValues) Get(key string) string {
	if vs := u[key]; len(vs) > 0 {
		return vs[0]
	}
	return ""
}

// ParseURL 解析原生url,获取URL Values
func ParseURL(r *http.Request) (URLValues, error) {
	var (
		values url.Values
		err    error
	)

	if r.URL != nil {
		values, err = url.ParseQuery(r.URL.RawQuery)
	}

	if values == nil {
		values = make(url.Values)
	}

	return URLValues(values), err
}
