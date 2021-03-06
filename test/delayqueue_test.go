package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/yasin-wu/delay-queue/delayqueue"
)

type JobActionSMS struct{}

func (JobActionSMS) ID() string {
	return "JobActionSMS"
}

func (JobActionSMS) Execute(args []interface{}) error {
	for _, arg := range args {
		if phoneNumber, ok := arg.(string); ok {
			fmt.Printf("sending sms to %s", phoneNumber)
		}
	}
	return nil
}

func TestDelayQueue(t *testing.T) {
	conf := &delayqueue.Config{
		Redis: &delayqueue.RedisConf{
			Host:     "192.168.131.135:6379",
			PassWord: "1qazxsw21201",
		},
	}
	dq := delayqueue.New(conf)
	err := dq.Register(JobActionSMS{})
	if err != nil {
		t.Errorf("register err:%v", err)
		return
	}
	dq.StartBackground()
	err = dq.AddJob(delayqueue.DelayJob{
		ID:        (&JobActionSMS{}).ID(),
		DelayTime: 2,
		Args:      []interface{}{"181****9331"},
	})
	if err != nil {
		t.Errorf("adddelay err:%v", err)
		return
	}
	time.Sleep(10 * time.Second)
}
