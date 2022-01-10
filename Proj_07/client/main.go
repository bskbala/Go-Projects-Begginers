package main

import (
	"log"
	"math/rand"
	"time"

	"tutorial-go-asynq/tasks"

	tasks "github.com/bskbala410@gmail.com/go-url-shortener/Proj_07/Tasks"
	"github.com/hibiken/asynq"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379",
	}
	client := asynq.NewClient(redisConnection)
	defer client.Close()

	for {
		userID := rand.Intn(1000) + 10

		delay := 2 * time.Minute

		task1 := tasks.NewWelcomeEmailTask(userID)
		task2 := tasks.NewReminderEmailTask(userID, time.Now().Add(delay))

		if _, err := client.Enqueue(
			task1,
			asynq.Queue("critical"),
		); err != nil {
			log.Fatal(err)

		}
		if _, err := client.Enqueue(
			task1,
			asynq.Queue("low"),
			asynq.ProcessIn(delay),
		); err != nil {
			log.Fatal(err)

		}
	}
}
