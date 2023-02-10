package schedule

import (
	"context"
	"fmt"

	"github.com/robfig/cron/v3"
)

type (
	Job      = cron.Job
	FuncJob  = cron.FuncJob
	Schedule = cron.Schedule
)

type Server struct {
	opts   []cron.Option
	c      *cron.Cron
	second bool
}

func (s Server) Start(_ context.Context) error {
	s.c.Start()
	return nil
}

func (s Server) Stop(_ context.Context) error {
	s.c.Stop()
	return nil
}

// NewServer 实例化cron扩展
func NewServer(opts ...ServerOption) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	s.c = cron.New(s.opts...)
	return s
}

// Job 设置job
func (s Server) Job(job Job) Spec {
	return NewSpec(s.AddJob, job)
}

// AddJob 根据cron表达式添加job
func (s Server) AddJob(spec interface{}, job Job) error {
	var err error

	switch expr := spec.(type) {
	case Schedule:
		_ = s.c.Schedule(expr, job)
		break
	case string:
		_, err = s.c.AddJob(expr, job)
	default:
		return fmt.Errorf("不支持此类型的解析表达式")
	}

	return err
}
