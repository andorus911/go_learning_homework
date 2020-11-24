package rabbit

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

func InitReminder(ctx context.Context, log *zap.Logger, rUsr, rPwd, rHost, rPort string) {
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "failed to register a consumer")

	lg.Info("Waiting for messages")
	for d := range msgs { // forever
		lg.Info("Received a message: " + string(d.Body))
	}
}