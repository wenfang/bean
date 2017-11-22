package conv

import (
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// StoStrings 将逗号分隔的字符串转换为字符串slice
func StoStrings(d string) []string {
	var s []string
	for _, v := range strings.Split(d, ",") {
		v = strings.TrimSpace(v)
		if v != "" {
			s = append(s, v)
		}
	}
	// 将结果排序
	sort.Strings(s)
	return s
}

// StoUint64s 将逗号分隔的字符串转换为uint64 slice
func StoUint64s(d string) ([]uint64, error) {
	var v []uint64
	for _, s := range strings.Split(d, ",") {
		s = strings.TrimSpace(s)
		if s != "" {
			i, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				return nil, errors.WithStack(err)
			}
			v = append(v, i)
		}
	}
	return v, nil
}

// RegularS 正则化逗号分隔的字符串，排序并去重
func RegularS(d string) string {
	// 转换为字符串数组
	v := StoStrings(d)
	if len(v) == 0 {
		return ""
	}
	// 排重
	var s []string
	last := v[0]
	s = append(s, last)
	for i := 1; i < len(v); i++ {
		if last != v[i] {
			last = v[i]
			s = append(s, last)
		}
	}
	return strings.Join(s, ",")
}
