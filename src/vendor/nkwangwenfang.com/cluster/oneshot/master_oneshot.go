package oneshot

import (
	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/master"
	"nkwangwenfang.com/log"
)

type masterOneshot struct {
	oneshot Oneshot
	master  master.Master
}

// WithMaster 包装oneshot使其具有master功能
func WithMaster(oneshot Oneshot, master master.Master) Oneshot {
	return &masterOneshot{oneshot: oneshot, master: master}
}

func (m *masterOneshot) Do(ctx context.Context) error {
	isMaster, err := m.master.CheckMaster()
	if err != nil {
		log.Error("check master failed", "error", err)
		return err
	}

	if isMaster {
		return m.oneshot.Do(ctx)
	}
	return nil
}
