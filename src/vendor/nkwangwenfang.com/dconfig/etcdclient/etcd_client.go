package etcdclient

import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"

	"nkwangwenfang.com/dconfig"
)

const prefix string = "dconfig"

type etcdClient struct {
	appName string
	keysAPI client.KeysAPI
}

// New 创建etcd客户端接口
func New(config Config) (dconfig.Client, error) {
	// 创建etcd client
	cfg := client.Config{
		Endpoints: config.Servers,
		Transport: client.DefaultTransport,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}

	ec := &etcdClient{
		appName: prefix + "/" + config.App,
		keysAPI: client.NewKeysAPI(c),
	}
	return ec, nil
}

func (ec *etcdClient) Get(ctx context.Context, key string) (string, error) {
	resp, err := ec.keysAPI.Get(ctx, ec.appName+"/"+key, nil)
	if err != nil {
		return "", err
	}
	if resp.Node == nil {
		return "", nil
	}
	return resp.Node.Value, nil
}
