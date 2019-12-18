package pumari

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type servicioRegistroDeEventosPumari struct {
	contextTimeout time.Duration
	channel        *amqp.Channel
	routingKey     string
}

func NuevoServicioRegistroDeEventosPumari(timeout time.Duration, rabbitMqHost string, rabbitMqPort string) controlasistencia.ServicioRegistroEventos {
	channel, routingKey := configureRabbitMq(rabbitMqHost, rabbitMqPort)
	return &servicioRegistroDeEventosPumari{
		contextTimeout: timeout,
		channel:        channel,
		routingKey:     routingKey,
	}
}

type MesanjeEvento struct {
	UserName string             `json:"user_name"`
	Tipo     modelos.TipoEvento `json:"tipo"`
	Fecha    time.Time          `json:"fecha"`
}

func (kq MesanjeEvento) Marshal() ([]byte, error) {
	bytes, err := json.Marshal(kq)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (s servicioRegistroDeEventosPumari) RegistrarEntrada(ctx context.Context, usuario modelos.Usuario, date time.Time) error {
	data := MesanjeEvento{
		UserName: usuario.ID,
		Tipo:     modelos.EntradaTipoEvento,
		Fecha:    date,
	}
	body, err := data.Marshal()
	if err != nil {
		return err
	}
	return s.send(body)
}

func (s servicioRegistroDeEventosPumari) RegistrarSalida(ctx context.Context, usuario modelos.Usuario, date time.Time) error {
	data := MesanjeEvento{
		UserName: usuario.ID,
		Tipo:     modelos.SalidaTipoEvento,
		Fecha:    date,
	}
	body, err := data.Marshal()
	if err != nil {
		return err
	}
	return s.send(body)
}

func configureRabbitMq(rabbitMqHost string, rabbitMqPort string) (*amqp.Channel, string) {
	connectionURL := fmt.Sprintf("amqp://guest:guest@%s:%s/", rabbitMqHost, rabbitMqPort)
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Panic(err)
	}
	//defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		log.Panic(err)
	}
	//defer ch.Close()
	routingKey := "kudos"
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

func (s servicioRegistroDeEventosPumari) send(body []byte) error {
	err := s.channel.Publish(
		"",
		s.routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		return err
	}
	return nil
}
