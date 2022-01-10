package tasks

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

func HandleWelcomeEmailTask(c context.Context, t *asynq.Task) error {
	id, err := t.payload.GetInt("user_id")
	if err != nil {
		return err
	}
	fmt.Printf("Send Welcome Email to the User ID %d \n,id")
	return nil
}
func HandleReminderEmailTask(c context.Context, t *asynq.Task) error {
	id, err := t.payload.GetInt("user_id")
	if err != nil {
		return err
	}
	time, err := t.Payload.GetString("sent_in")
	if err != nil {
		return err
	}
	fmt.Printf("Send Reminder Email to the User ID %d \n,id")
	fmt.Printf("Reason :time is up (%v)\n", time)
	return nil
}
