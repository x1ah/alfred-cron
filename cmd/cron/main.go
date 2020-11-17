package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	parser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	s, err := parser.Parse("1 * * * *")
	if err != nil {
		panic(err)
	}
	curr := time.Now()
	for i := 0; i <= 10; i++ {
		curr = s.Next(curr)
		fmt.Println(curr)
	}
}
