package schedule

import (
	"fmt"
	"math"
	"strings"
)

type Spec interface {
	Cron(expression interface{}) error
	EverySeconds(seconds ...int) error
	EveryMinute() error
	AfterEveryMinute() error
	EveryMinutes(minutes int) error
	AfterEveryMinutes(minutes int) error
	EveryFiveMinutes() error
	AfterEveryFiveMinutes() error
	EveryTenMinutes() error
	AfterEveryTenMinutes() error
	EveryFifteenMinutes() error
	AfterEveryFifteenMinutes() error
	EveryThirtyMinutes() error
	AfterEveryThirtyMinutes() error
	Hourly() error
	HourlyAt(offset int) error
	Daily() error
	DailyAt(t string) error
	Weekly() error
	Monthly() error
}

const (
	// Set the top bit if a star was included in the expression.
	starBit = 1 << 63
)

type bounds struct {
	min, max uint
	names    map[string]uint
}

// getBits sets all bits in the range [min, max], modulo the given step size.
func getBits(min, max, step uint) uint64 {
	var bits uint64

	// If step is 1, use shifts.
	if step == 1 {
		return ^(math.MaxUint64 << (max + 1)) & (math.MaxUint64 << min)
	}

	// Else, use a simple loop.
	for i := min; i <= max; i += step {
		bits |= 1 << i
	}
	return bits
}

// all returns all bits within the given bounds.  (plus the star bit)
func all(r bounds) uint64 {
	return getBits(r.min, r.max, 1) | starBit
}

var _ Spec = new(spec)

type JobRunner func(spec interface{}, job Job) error

type spec struct {
	job    Job
	runner JobRunner
}

// NewSpec 实例化cron表达式
func NewSpec(runner JobRunner, job Job) *spec {
	return &spec{
		job:    job,
		runner: runner,
	}
}

// Cron 自定义cron表达式运行job
func (s spec) Cron(spec interface{}) error {
	return s.runner(spec, s.job)
}

// EverySeconds 每秒运行job
func (s spec) EverySeconds(seconds ...int) error {
	defaultSeconds := 1
	if len(seconds) > 0 {
		defaultSeconds = seconds[0]
	}
	return s.Cron(fmt.Sprintf("@every %ds", defaultSeconds))
}

// EveryMinute 每分钟运行job
func (s spec) EveryMinute() error {
	return s.EveryMinutes(1)
}

// AfterEveryMinute 每分钟后运行job
func (s spec) AfterEveryMinute() error {
	return s.Cron("@every 1m")
}

// EveryMinutes 每多少分钟运行job
func (s spec) EveryMinutes(minutes int) error {
	return s.Cron(fmt.Sprintf("0 0/%d * * * *", minutes))
}

// AfterEveryMinutes 每多少分钟后运行job
func (s spec) AfterEveryMinutes(minutes int) error {
	return s.Cron(fmt.Sprintf("@every %dm", minutes))
}

// EveryFiveMinutes 每五分钟运行job
func (s spec) EveryFiveMinutes() error {
	return s.EveryMinutes(5)
}

// AfterEveryFiveMinutes 每五分钟后运行job
func (s spec) AfterEveryFiveMinutes() error {
	return s.AfterEveryMinutes(5)
}

// EveryTenMinutes 每十分钟运行job
func (s spec) EveryTenMinutes() error {
	return s.EveryMinutes(10)
}

// AfterEveryTenMinutes 每十分钟后运行job
func (s spec) AfterEveryTenMinutes() error {
	return s.AfterEveryMinutes(10)
}

// EveryFifteenMinutes 每十五分钟运行job
func (s spec) EveryFifteenMinutes() error {
	return s.EveryMinutes(15)
}

// AfterEveryFifteenMinutes 每十五分钟后运行job
func (s spec) AfterEveryFifteenMinutes() error {
	return s.AfterEveryMinutes(15)
}

// EveryThirtyMinutes 每三十分钟运行job
func (s spec) EveryThirtyMinutes() error {
	return s.EveryMinutes(30)
}

// AfterEveryThirtyMinutes 每三十分钟后运行job
func (s spec) AfterEveryThirtyMinutes() error {
	return s.AfterEveryMinutes(30)
}

// Hourly 每小时运行job
func (s spec) Hourly() error {
	return s.Cron("@hourly")
}

// HourlyAt 每小时的某分钟运行job
func (s spec) HourlyAt(offset int) error {
	return s.Cron(fmt.Sprintf("@every 1h%dm", offset))
}

// Daily 每天运行job
func (s spec) Daily() error {
	return s.Cron("@daily")
}

// DailyAt 每天某时某分运行job
// DailyAt("12:21") 每天12点21分钟运行
func (s spec) DailyAt(t string) error {
	tt := strings.Split(t, ":")
	if len(tt) < 2 {
		tt = append(tt, "00")
	}
	return s.Cron(fmt.Sprintf("%s %s * * *", tt[0], tt[1]))
}

// Weekly 每周运行job
func (s spec) Weekly() error {
	return s.Cron("@weekly")
}

// Monthly 每月运行job
func (s spec) Monthly() error {
	return s.Cron("@monthly")
}
