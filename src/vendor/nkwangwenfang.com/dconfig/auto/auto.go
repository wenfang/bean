package auto

import (
	"sync"

	"golang.org/x/net/context"

	"nkwangwenfang.com/dconfig"
	"nkwangwenfang.com/log"
)

// AutoConfig 自动配置结构
type Auto struct {
	sync.RWMutex
	dconfig.Client
	values    map[string]string
	observers map[string][]Observer
}

// New 创建AutoConfig结构
func New(c dconfig.Client) *Auto {
	ac := &Auto{
		Client:    c,
		values:    make(map[string]string),
		observers: make(map[string][]Observer),
	}
	return ac
}

// Register 注册对应key的Observer
func (ac *Auto) Register(key string, value string, observer Observer) {
	ac.Lock()
	defer ac.Unlock()

	// 加入observers监控列表
	ac.observers[key] = append(ac.observers[key], observer)
	// 设置初值
	ac.values[key] = value
}

func (ac *Auto) Do(ctx context.Context) error {
	ac.RLock()
	defer ac.RUnlock()

	for key, observers := range ac.observers {
		value, err := ac.Get(ctx, key)
		if err != nil {
			log.Error("client get key error", "key", key, "error", err)
			continue
		}
		if value != ac.values[key] {
			ac.values[key] = value
			for _, observer := range observers {
				observer.ConfigUpdate(value)
			}
		}
	}
	return nil
}
