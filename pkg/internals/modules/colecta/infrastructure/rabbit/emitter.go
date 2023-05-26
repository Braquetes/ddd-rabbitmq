package rabbit

import (
	"fmt"

	"github.com/streadway/amqp"
)

func Prod(message string, key string, exchange string, types string) error {
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

	// args := amqp.Table{
	// 	"x-message-ttl": 5000, // Tiempo de vida del mensaje en milisegundos
	// 	// "x-dead-letter-exchange": "dead-letter", // Exchange al que se redirigirán los mensajes no entregables
	// }

	// declaring queue with its properties over the the channel opened
	queue, err := channel.QueueDeclare(
		"testing", // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		panic(err)
	}

	// err = channel.ExchangeDeclare(
	// 	"",    // nombre del exchange
	// 	types, // tipo de exchange
	// 	true,  // duradero
	// 	false, // no autodelete
	// 	false, // sin interno
	// 	false, // sin espera de confirmación
	// 	nil,   // argumentos adicionales
	// )
	// if err != nil {
	// 	return err
	// }

	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		return err
	}

	fmt.Println("Queue:", queue)
	fmt.Printf("Enviado a: %s\n", key)
	return nil
}
