package dconfig

import (
	"golang.org/x/net/context"
)

// Load 利用client装载对应key的配置信息
func Load(ctx context.Context, client Client, key string) (string, error) {
	value, err := client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return value, nil
}
