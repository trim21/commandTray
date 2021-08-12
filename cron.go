package main

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

func buildCron(config T) *cron.Cron {
	c := cron.New(
		cron.WithParser(cron.NewParser(cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor)),
		cron.WithLocation(tz),
	)

	for _, c2 := range config.Cron {
		if c2.check() == nil {
			_, err := c.AddFunc(c2.Spec, c2.execute)
			if err != nil {
				logrus.Errorln(err)
			}
		}
	}

	return c
}
