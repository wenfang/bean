package dconfig

import (
	"context"
)

// Client 接口，获取配置信息
type Client interface {
	Get(context.Context, string) (string, error)
}

// Load 利用client装载对应key的配置信息
func Load(ctx context.Context, client Client, key string) (string, error) {
	value, err := client.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return value, nil
}
