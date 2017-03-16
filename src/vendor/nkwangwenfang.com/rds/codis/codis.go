package codis

import (
	"math/rand"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/pkg/errors"
)

var (
	// ErrNoConn 无连接错误
	ErrNoConn = errors.New("no redis connection")
)

// Codis 结构
type Codis []*redis.Pool

// New 创建访问Codis的连接池
func New(config Config) Codis {
	co := make(Codis, len(config.Addrs))
	for idx, addr := range config.Addrs {
		co[idx] = &redis.Pool{
			MaxIdle:     config.MaxIdle,
			MaxActive:   config.MaxActive,
			IdleTimeout: config.IdleTimeout * time.Millisecond,
			Wait:        config.Wait,
			Dial: func() (redis.Conn, error) {
				c, err := redis.DialTimeout(
					"tcp",
					addr,
					config.ConnectTimeout*time.Millisecond,
					config.ReadTimeout*time.Millisecond,
					config.WriteTimeout*time.Millisecond,
				)
				if err != nil {
					return nil, errors.Wrap(err, "redis dial error")
				}
				return c, nil
			},
		}
	}
	return co
}

func (co Codis) getConn() redis.Conn {
	cnt := len(co)
	if cnt == 0 {
		return nil
	}
	return co[rand.Intn(cnt)].Get()
}

// DoCmd 执行redis命令
func (co Codis) DoCmd(cmd string, args ...interface{}) (interface{}, error) {
	c := co.getConn()
	if c == nil {
		return nil, ErrNoConn
	}
	defer c.Close()

	return c.Do(cmd, args...)
}

// DoScript 执行redis脚本
func (co Codis) DoScript(script *redis.Script, args ...interface{}) (interface{}, error) {
	c := co.getConn()
	if c == nil {
		return nil, ErrNoConn
	}
	defer c.Close()

	return script.Do(c, args...)
}

// Close 关闭redis连接池
func (co Codis) Close() {
	for _, p := range co {
		p.Close()
	}
}
