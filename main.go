package main

import (
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
	crondesc "github.com/lnquy/cron"
	"github.com/robfig/cron/v3"
)

// AlfredCron alfred-cron struct
type AlfredCron struct {
	wf *aw.Workflow
	// how many times
	repeated   int
	cronParser cron.Parser
	descr      *crondesc.ExpressionDescriptor
}

// NewAlfredCron create new AlfredCron instance
func NewAlfredCron() *AlfredCron {
	descr, _ := crondesc.NewDescriptor()
	return &AlfredCron{
		wf:         aw.New(),
		repeated:   7,
		cronParser: cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
		descr:      descr,
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
		if description, err := ac.descr.ToDescription(spec, crondesc.Locale_en); err == nil {
			ac.wf.NewItem(description)
		}
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
