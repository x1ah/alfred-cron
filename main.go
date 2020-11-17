package main

import (
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/robfig/cron/v3"
)

// AlfredCron alfred-cron struct
type AlfredCron struct {
	wf *aw.Workflow
	// how many times
	repeated   int
	cronParser cron.Parser
}

// NewAlfredCron create new AlfredCron instance
func NewAlfredCron() *AlfredCron {
	return &AlfredCron{
		wf:         aw.New(),
		repeated:   7,
		cronParser: cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
	}
}

// Run run alfred-cron workflow
func (ac *AlfredCron) Run() {
	spec := strings.Join(ac.wf.Args(), "")
	if spec == "" {
		ac.wf.NewItem("no crontab expression found.").Subtitle("Try input a crontab expression?")
		ac.wf.SendFeedback()
		return
	}

	expr, err := ac.cronParser.Parse(spec)
	if err != nil {
		ac.wf.NewWarningItem("Invalid expression", err.Error())
	} else {
		curr := time.Now()
		for i := 0; i < ac.repeated; i++ {
			curr = expr.Next(curr)
			ac.wf.NewItem(curr.String())
		}
	}
	ac.wf.SendFeedback()
}

func main() {
	alfredCron := NewAlfredCron()
	alfredCron.Run()
}
