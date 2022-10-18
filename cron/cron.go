package cron

import (
	"github.com/BlueSky1405/project-kit/log"
	"github.com/pkg/errors"
	"github.com/robfig/cron/v3"
)

type Cron struct {
	logger log.Logger
	cron   *cron.Cron
	// 定时器任务与时间表达式映射
	entries map[string]string
}

func NewCron(logger log.Logger) *Cron {
	return &Cron{
		logger:  logger,
		entries: make(map[string]string),
		cron:    cron.New(cron.WithSeconds()),
	}
}

func (c *Cron) Start() {
	c.logger.InfoW("Cron is starting...")
	if c.entries == nil || len(c.entries) == 0 {
		c.logger.WarnW("Cron entries is nil")
		return
	}

	for name, expression := range c.entries {
		c.logger.InfoW("Cron register", "job_name", name, "expression", expression)
	}

	c.cron.Start()
}

func (c *Cron) Stop() {
	c.logger.InfoW("Cron is stopping...")
	c.cron.Stop()
}

// AddFunc jobName: 任务名称, spec: s m h d M w, cmd: 处理函数
func (c *Cron) AddFunc(jobName, spec string, cmd func()) error {
	_, err := c.cron.AddFunc(spec, cmd)
	if err != nil {
		return errors.Wrapf(err, "Cron AddFunc fail")
	}

	// 不做判断了，后一个覆盖前一个
	c.entries[jobName] = spec

	return nil
}
