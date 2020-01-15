package log

import (
	"context"
	"pumari/modelos"
)

type CasoDeUso interface {
	RegistroEvento(ctx context.Context, mensaje modelos.MensanjeEvento) error
}

type ServicioLog interface {
	RegistrarEvento(ctx context.Context, mensaje modelos.MensanjeEvento) error
}
