package endpoint

import (
	"context"
)

// Endpoint 代表单个RPC方法
type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

// Middleware 一层层封装Endpoint的中间件
type Middleware func(Endpoint) Endpoint

// Chain 创建中间件的辅助函数
func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Endpoint) Endpoint {
		for i := len(others) - 1; i >= 0; i-- {
			next = others[i](next)
		}
		return outer(next)
	}
}
