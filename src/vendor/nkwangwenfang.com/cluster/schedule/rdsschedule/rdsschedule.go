package rdsschedule

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/garyburd/redigo/redis"

	"nkwangwenfang.com/cluster/schedule"
	"nkwangwenfang.com/log"
	"nkwangwenfang.com/rds"
)

const scheduleKey string = "%s_schedule"

type rdsSchedule struct {
	r   rds.Rdser
	key string
}

func New(r rds.Rdser, name string) schedule.Scheduler {
	return &rdsSchedule{
		r:   r,
		key: fmt.Sprintf(scheduleKey, name),
	}
}

func (rs *rdsSchedule) GetInfo() (schedule.Info, error) {
	reply, err := rs.r.DoCmd("GET", rs.key)
	if err != nil {
		log.Error("redis get info error", "key", rs.key, "error", err)
		return nil, err
	}
	msg, err := redis.String(reply, err)
	if err != nil {
		log.Error("redis info is nil", "key", rs.key)
		return nil, errors.New("redis info is nil")
	}
	// 解析为Info信息
	var info schedule.Info
	if err := json.Unmarshal([]byte(msg), &info); err != nil {
		log.Error("redis info unmarshal error", "scheduleKey", rs.key, "msg", msg, "error", err)
		return nil, err
	}
	return info, nil
}

func (rs *rdsSchedule) SetInfo(info schedule.Info) error {
	msg, err := json.Marshal(info)
	if err != nil {
		log.Error("redis info marshal error", "info", info, "error", err)
		return err
	}
	if _, err = rs.r.DoCmd("SET", rs.key, string(msg)); err != nil {
		log.Error("redis set info error", "scheduleKey", rs.key, "msg", string(msg), "error", err)
		return err
	}
	return nil
}
