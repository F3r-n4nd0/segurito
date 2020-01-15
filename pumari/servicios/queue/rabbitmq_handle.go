package queue

import (
	"context"
	"log"
	log2 "pumari/log"
	"pumari/modelos"

	"github.com/streadway/amqp"
)

func NewQueueHttpHandler(channel *amqp.Channel, routingKey string, casoDeUsoLog log2.CasoDeUso) {

	msgs, err := channel.Consume(
		routingKey,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			mensaje, err := modelos.UnMarshal(d.Body)
			if err != nil {
				log.Print(err)
				continue
			}
			err = casoDeUsoLog.RegistroEvento(context.Background(), *mensaje)
			if err != nil {
				log.Print("ERROR: " + err.Error())
			}
		}
	}()

	<-forever
}
