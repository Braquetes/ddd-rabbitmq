package main

import (
	"fmt"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Go-RabbitMQ")

	connection, err := amqp.Dial("amqp://braquetes:1902@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer connection.Close()

	fmt.Println("Conectado")

	channel, err := connection.Channel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	// fmt.Println("Create")
	// // Crear una cola
	// queue, err := channel.QueueDeclare(
	// 	"",    // Nombre de la cola
	// 	false, // durable
	// 	false, // delete when unused
	// 	false, // exclusive
	// 	false, // no-wait
	// 	nil,   // arguments
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("QueueBind", queue.Name, queue.Consumers, queue.Messages)

	// err = channel.QueueBind(
	// 	queue.Name,    // Nombre de la cola
	// 	"xdxd",        // Routing key
	// 	"my_exchange", // Nombre del exchange
	// 	false,         // No-wait
	// 	nil,           // Arguments
	// )
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Consume")

	// docker exec -it go-rabbitmq
	msgs, err := channel.Consume(
		"testing", // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       //args
	)
	if err != nil {
		panic(err)
	}

	forever := make(chan bool)
	i := 0
	go func() {
		for msg := range msgs {

			// fmt.Printf("%+v\n", msg)

			// // Acknowledge (ACK) el mensaje
			// err := msg.Ack(false)
			// if err != nil {
			// 	// Manejar el error
			// 	panic(err)
			// }

			fmt.Printf("Entrada %d: %s - %v \n", i, msg.Body, msg)
			i++
		}
	}()

	fmt.Println("Escuchando ...")
	<-forever
}
