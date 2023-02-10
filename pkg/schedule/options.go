package schedule

import (
	"time"

	"github.com/robfig/cron/v3"
)

type (
	Parser     = cron.ScheduleParser
	JobWrapper = cron.JobWrapper
	Logger     = cron.Logger
)

const (
	Second         = cron.Second
	SecondOptional = cron.SecondOptional
	Minute         = cron.Minute
	Hour           = cron.Hour
	Dom            = cron.Dom
	Month          = cron.Month
	Dow            = cron.Dow
	DowOptional    = cron.DowOptional
	Descriptor     = cron.Descriptor
)

var (
	NewParser = cron.NewParser
)

type ServerOption func(o *Server)

func WithLocation(loc *time.Location) ServerOption {
	return func(o *Server) {
		o.opts = append(o.opts, cron.WithLocation(loc))
	}
}

func WithChain(wrappers ...JobWrapper) ServerOption {
	return func(o *Server) {
		o.opts = append(o.opts, cron.WithChain(wrappers...))
	}
}

func WithSeconds() ServerOption {
	return func(o *Server) {
		o.opts = append(o.opts, cron.WithSeconds())
		o.second = true
	}
}

func WithParser(parser Parser) ServerOption {
	return func(o *Server) {
		o.opts = append(o.opts, cron.WithParser(parser))
	}
}

func WithLogger(logger Logger) ServerOption {
	return func(o *Server) {
		o.opts = append(o.opts, cron.WithLogger(logger))
	}
}
