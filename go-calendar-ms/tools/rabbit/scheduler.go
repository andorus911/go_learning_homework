package rabbit

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"go_learning_homework/go-calendar-ms/internal/domain/models"
	"go_learning_homework/go-calendar-ms/tools/postgres"
	"time"
)

var lg zap.Logger

func failOnError(err error, msg string) {
	if err != nil {
		lg.Fatal(msg + err.Error())
	}
}

func InitScheduler(ctx context.Context, log *zap.Logger, db postgres.DB, rUsr, rPwd, rHost, rPort string) {
	lg = *log

	dialUrl := fmt.Sprintf("amqp://%v:%v@%v:%v/", rUsr, rPwd, rHost, rPort)
	conn, err := amqp.Dial(dialUrl)
	failOnError(err, "failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"scheduler",
		false,
		false,
		false,
		false,
		nil,
		)
	failOnError(err, "failed to declare a queue")

	toRemind := make(chan models.Event)
	go func(ctx context.Context) {
		for {
			checkDB(ctx, toRemind, db)
			time.Sleep(time.Second)
		}
	}(ctx)

	for e := range toRemind {
		body, err := json.Marshal(e)
		failOnError(err, "failed to marshal an event")

		err = ch.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        []byte(body),
			})
		lg.Info(" [x] Sent " + string(body))
		failOnError(err, "failed to publish a message")
	}
	lg.Info("THE END")
}

func checkDB(ctx context.Context, toRemind chan<- models.Event, db postgres.DB) {
	events, err := db.GetAllEventsFromNow(ctx) // todo query not for all events
	failOnError(err, "failed to get events from db")

	for _, e := range events {
		now := time.Now()
		sub := e.StartTime.Sub(now)
		min := 5 * time.Minute - time.Second
		max := 5 * time.Minute
		if sub > min && sub < max {
			toRemind <- e
		}
	}
}