package rdskoffset

import (
	"fmt"

	"github.com/garyburd/redigo/redis"

	"nkwangwenfang.com/kafka/koffset"
	"nkwangwenfang.com/rds"
)

const (
	poKey string = "%s_partition_offset_%d"
)

type rdsKOffset struct {
	rds.Rdser
	app string
}

// New 创建基于redis的Offsetter
func New(app string, r rds.Rdser) koffset.KOffsetter {
	return &rdsKOffset{app: app, Rdser: r}
}

// Get 获取partition对应的offset
func (rko *rdsKOffset) Get(partitionID int32) (int64, error) {
	key := fmt.Sprintf(poKey, rko.app, partitionID)
	off, err := redis.Int64(rko.DoCmd("GET", key))
	if err != nil {
		return 0, err
	}
	return off, nil
}

// Set 记录partition对应的offset
func (rko *rdsKOffset) Set(partitionID int32, off int64) error {
	key := fmt.Sprintf(poKey, rko.app, partitionID)
	if _, err := rko.DoCmd("SET", key, off); err != nil {
		return err
	}
	return nil
}
