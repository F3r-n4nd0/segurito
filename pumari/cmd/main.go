package main

import (
	"fmt"
	"log"
	"os"
	log2 "pumari/log"
	mesa "pumari/servicios/pumari"
	"pumari/servicios/queue"

	"github.com/streadway/amqp"
)

func main() {

	rabbitMqHost := getenv("RABBIT_MQ_HOST", "rabbit_mq")
	rabbitMqPort := getenv("RABBIT_MQ_PORT", "5672")
	pumariPath := getenv("PUMARI_PATH", "./mocks/")

	rabbitMqChannel, routingKey := configureRabbitMq(rabbitMqHost, rabbitMqPort)
	serviceLog := mesa.NuevoServicioDeLogPumari(2, pumariPath)
	userCaseLog := log2.NuevoCasoDeUsoLog(20, serviceLog)

	log.Print("Servicio Pumari iniciado")
	forever := make(chan bool)
	queue.NewQueueHttpHandler(rabbitMqChannel, routingKey, userCaseLog)
	<-forever

}

func configureRabbitMq(rabbitMqHost string, rabbitMqPort string) (*amqp.Channel, string) {

	connectionURL := fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitMqHost, rabbitMqPort)
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	routingKey := "segurito"
	q, err := ch.QueueDeclare(
		routingKey,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Panic(err)
	}
	return ch, q.Name
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
