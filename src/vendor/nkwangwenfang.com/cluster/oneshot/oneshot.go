package oneshot

import (
	"golang.org/x/net/context"
)

// Oneshot 每次Do返回后任务创建完成，任务会自动终止
type Oneshot interface {
	Do(context.Context) error
}
