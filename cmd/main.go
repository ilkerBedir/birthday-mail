package main

import (
	"context"
	"fmt"
	"log"
	"project/birthday-mail/internal/config"
	"project/birthday-mail/internal/database"
	"project/birthday-mail/internal/mail"
	"time"

	_ "github.com/lib/pq"
	"github.com/reugn/go-quartz/job"
	"github.com/reugn/go-quartz/quartz"
)

func main() {
	log.Println("Starting...")
	config.GetConfig()
	defer database.Close()
	sendMailToUsers()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sched := quartz.NewStdScheduler()
	sched.Start(ctx)
	cronTrigger, _ := quartz.NewCronTrigger("0 1 * * * *")

	functionJob := job.NewFunctionJob(func(_ context.Context) (int, error) {
		log.Println("Schedule started...")
		sendMailToUsers()
		return 0, nil
	})

	sched.ScheduleJob(quartz.NewJobDetail(functionJob, quartz.NewJobKey("functionJob")),
		cronTrigger)
	sched.Wait(ctx)
}

func sendMailToUsers() {
	now := time.Now()
	now2 := fmt.Sprintf("%s", now.Format("2006-01-02"))
	fmt.Println(now2)
	users, err := database.SelectUsersByDate(now2)
	if err != nil {
		log.Fatalln(err)
	}
	for _, user := range users {
		mail.SendMail(user)
	}
}
