package handler

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// URLParam 代表Url参数值
type URLParam url.Values

// GetString 获取名称为key的第一个值string类型
func (u URLParam) GetString(key string) string {
	vs := u[key]
	if len(vs) == 0 {
		return ""
	}
	return vs[0]
}

// GetStrings 获取名称为key的第一个值的string列表类型
func (u URLParam) GetStrings(key string) []string {
	vs := u[key]
	if len(vs) == 0 {
		return nil
	}

	var rvs []string
	for _, v := range strings.Split(vs[0], ",") {
		if rv := strings.TrimSpace(v); rv != "" {
			rvs = append(rvs, rv)
		}
	}
	return rvs
}

// GetInt64 获取名称为key的第一个值int64类型
func (u URLParam) GetInt64(key string) (int64, error) {
	vs := u[key]
	if len(vs) == 0 {
		return 0, errors.Errorf("parameter %s not found", key)
	}

	v, err := strconv.ParseInt(vs[0], 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "parameter %s to int error", key)
	}
	return v, nil
}

// GetInt64s 获取名称为key的第一个值的int64列表类型
func (u URLParam) GetInt64s(key string) ([]int64, error) {
	vs := u[key]
	if len(vs) == 0 {
		return nil, errors.Errorf("parameter %s not found", key)
	}

	var rvs []int64
	for _, v := range strings.Split(vs[0], ",") {
		rv, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parameter %s to int slice error", key)
		}
		rvs = append(rvs, rv)
	}
	return rvs, nil
}

// GetUint64 获取名称为key的第一个值uint64类型
func (u URLParam) GetUint64(key string) (uint64, error) {
	vs := u[key]
	if len(vs) == 0 {
		return 0, errors.Errorf("parameter %s not found", key)
	}

	v, err := strconv.ParseUint(vs[0], 10, 64)
	if err != nil {
		return 0, errors.Wrapf(err, "parameter %s to uint error", key)
	}
	return v, nil
}

// GetUint64s 获取名称为key的第一个值uint64列表类型
func (u URLParam) GetUint64s(key string) ([]uint64, error) {
	vs := u[key]
	if len(vs) == 0 {
		return nil, errors.Errorf("parameter %s not found", key)
	}

	var rvs []uint64
	for _, v := range strings.Split(vs[0], ",") {
		rv, err := strconv.ParseUint(strings.TrimSpace(v), 10, 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parameter %s to uint slice error", key)
		}
		rvs = append(rvs, rv)
	}
	return rvs, nil
}

// GetFloat64 获取名称为key的第一个值float64类型
func (u URLParam) GetFloat64(key string) (float64, error) {
	vs := u[key]
	if len(vs) == 0 {
		return 0, errors.Errorf("parameter %s not found", key)
	}

	v, err := strconv.ParseFloat(vs[0], 64)
	if err != nil {
		return 0, errors.Wrapf(err, "parameter %s to float error", key)
	}
	return v, nil
}

// GetFloat64s 获取名称为key的第一个值float64列表类型
func (u URLParam) GetFloat64s(key string) ([]float64, error) {
	vs := u[key]
	if len(vs) == 0 {
		return nil, errors.Errorf("parameter %s not found", key)
	}

	var rvs []float64
	for _, v := range strings.Split(vs[0], ",") {
		rv, err := strconv.ParseFloat(strings.TrimSpace(v), 64)
		if err != nil {
			return nil, errors.Wrapf(err, "parameter %s to float slice error", key)
		}
		rvs = append(rvs, rv)
	}
	return rvs, nil
}

// ParseURL 解析原生url,获取URL Values
func ParseURL(r *http.Request) (URLParam, error) {
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
	return URLParam(values), err
}
