package main

import (
	"log"

	"tutorial-go-asynq/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	redisConnection := asynq.RedisClientOpt{

		Addr: "localhost:6379",
	}
	worker := asynq.NewServer(redisConnection, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"critical": 6,
			"default":  3,
			"low":      1,
		},
	})

	mux := asynq.NewServeMux()
	mux.Handle(
		tasks.TypeWelcomeEmail,
		tasks.HandleWelcomeEmailTask,
	)
	mux.HandleFunc(
		tasks.TypeReminderEamil,
		tasks.HandleReminderEmailTask,
	)
	if err := worker.Run(mux); err != nil {
		log.Fatal(err)
	}
}
