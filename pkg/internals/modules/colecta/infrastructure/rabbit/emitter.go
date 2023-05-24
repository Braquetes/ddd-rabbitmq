package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Prod(message string, key string) error {
	fmt.Printf("Go-RabbitMQ")

	connection, err := amqp.Dial("amqp://braquetes:1902@localhost:5672/")
	if err != nil {
		return err
	}
	defer connection.Close()

	fmt.Println("Conectado")

	channel, err := connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	// queue, err := channel.QueueDeclare(
	// 	"testing", // name
	// 	false,     // durable
	// 	false,     // auto delete
	// 	false,     // exclusive
	// 	false,     // no wait
	// 	nil,       // args
	// )
	// if err != nil {
	// 	return err
	// }

	err = channel.ExchangeDeclare(
		"my_exchange", // nombre del exchange
		"fanout",      // tipo de exchange
		true,          // duradero
		false,         // no autodelete
		false,         // sin interno
		false,         // sin espera de confirmación
		nil,           // argumentos adicionales
	)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"my_exchange", // exchange
		key,           // key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
			// UserId:      "1",
		},
	)

	if err != nil {
		return err
	}

	// fmt.Println("Queue:", queue)
	fmt.Printf("Enviado a: %s\n", key)
	return nil
}
