package tasks

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeWelcomeEmail = "email:welcome"

	TypeReminderEamil = "email:reminder"
)

func NewWelcomeEmailTask(id int) *asynq.Task {
	payload := map[string]interface{}{
		"userr_id": id,
	}
	return asynq.NewTask(TypeWelcomeEmail, payload)
}

func NewReminderEmailTask(id int, ts time.Time) *asynq.Task {
	payload := map[string]interface{}{
		"userr_id": id,
		"sent_in":  ts.String(),
	}
	return asynq.NewTask(TypeReminderEmail, payload)
}
