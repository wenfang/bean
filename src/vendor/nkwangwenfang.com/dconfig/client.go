package dconfig

import "golang.org/x/net/context"

// Client 接口，获取配置信息
type Client interface {
	Get(context.Context, string) (string, error)
}
