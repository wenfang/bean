package rdsdworker

import (
	"testing"
	"time"

	"golang.org/x/net/context"

	"nkwangwenfang.com/cluster/task"
	"nkwangwenfang.com/log"
	"nkwangwenfang.com/rds/codis"
)

var codisConfig = codis.Config{
	Addrs:          []string{"127.0.0.1:6379"},
	MaxIdle:        20,
	MaxActive:      50,
	ConnectTimeout: 3000,
	ReadTimeout:    3000,
	WriteTimeout:   3000,
	IdleTimeout:    30000,
	Wait:           true,
}

type testTask struct {
	ID int
}

func (tt *testTask) Start(c context.Context) (context.CancelFunc, chan struct{}, error) {
	ctx, cancel := context.WithCancel(c)
	done := make(chan struct{})

	go func() {
		defer close(done)

		tick := time.Tick(3e9)
		for {
			select {
			case <-tick:
			case <-ctx.Done():
				log.Info("testTask Done", "ID", tt.ID)
				return
			}

		}
	}()
	return cancel, done, nil
}

type testFactory struct{}

func (tf *testFactory) Name() string {
	return "testFactory"
}

func (tf *testFactory) CreateTask(id int) task.Task {
	return &testTask{ID: id}
}

func (tf *testFactory) AllTaskIDs() []int {
	return []int{1, 2, 3, 4, 5}
}

func TestBundle(t *testing.T) {
	r := codis.New(codisConfig)
	defer r.Close()

	num := 3
	bundle := make([]*Bundle, num)

	for i := 0; i < num; i++ {
		bundle[i] = New(r, "bundle", &testFactory{})
		if err := bundle[i].Start(context.Background()); err != nil {
			t.Fatal(err)
		}
	}

	time.Sleep(30e9)

	for i := 0; i < num; i++ {
		time.Sleep(20e9)
		bundle[i].Stop()
	}
}
