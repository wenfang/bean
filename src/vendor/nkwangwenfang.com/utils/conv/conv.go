package conv

import (
	"reflect"
	"strconv"

	"github.com/pkg/errors"
)

// ItoInt 接口类型转换为int
func ItoInt(d interface{}) (int, error) {
	switch v := d.(type) {
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case int:
		return int(v), nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return int(v), nil
	case uint64:
		return int(v), nil
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return int(i), nil
	}
	return 0, errors.Errorf("invalid type %s", reflect.TypeOf(d))
}

// ItoInt64 接口类型转换为int64
func ItoInt64(d interface{}) (int64, error) {
	switch v := d.(type) {
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return int64(v), nil
	case uint:
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return int64(v), nil
	case string:
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return i, nil
	}
	return 0, errors.Errorf("invalid type %s", reflect.TypeOf(d))
}

// ItoUint 接口类型转换为uint
func ItoUint(d interface{}) (uint, error) {
	switch v := d.(type) {
	case float32:
		return uint(v), nil
	case float64:
		return uint(v), nil
	case int:
		return uint(v), nil
	case int8:
		return uint(v), nil
	case int16:
		return uint(v), nil
	case int32:
		return uint(v), nil
	case int64:
		return uint(v), nil
	case uint:
		return uint(v), nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		return uint(v), nil
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return uint(u), nil
	}
	return 0, errors.Errorf("invalid type %s", reflect.TypeOf(d))
}

// ItoUint64 接口类型转换为uint64
func ItoUint64(d interface{}) (uint64, error) {
	switch v := d.(type) {
	case float32:
		return uint64(v), nil
	case float64:
		return uint64(v), nil
	case int:
		return uint64(v), nil
	case int8:
		return uint64(v), nil
	case int16:
		return uint64(v), nil
	case int32:
		return uint64(v), nil
	case int64:
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return uint64(v), nil
	case string:
		u, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return 0, err
		}
		return u, nil
	}
	return 0, errors.Errorf("invalid type %s", reflect.TypeOf(d))
}

// ItoFloat64 接口类型转换为float64
func ItoFloat64(d interface{}) (float64, error) {
	switch v := d.(type) {
	case float32:
		return float64(v), nil
	case float64:
		return float64(v), nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case string:
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return 0.0, err
		}
		return f, nil
	}
	return 0, errors.Errorf("invalid type %s", reflect.TypeOf(d))
}

// ItoString 接口类型转换为string
func ItoString(d interface{}) (string, error) {
	switch v := d.(type) {
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 64), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case int:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int32:
		return strconv.FormatInt(int64(v), 10), nil
	case int64:
		return strconv.FormatInt(int64(v), 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(v), 10), nil
	case string:
		return v, nil
	}
	return "", errors.Errorf("invalid type %s", reflect.TypeOf(d))
}
