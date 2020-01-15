package log

import (
	"context"
	"pumari/modelos"
	"time"
)

type logCasoDeUso struct {
	contextTimeout time.Duration
	servicioLog    ServicioLog
}

func NuevoCasoDeUsoLog(timeout time.Duration, servicioLog ServicioLog) CasoDeUso {
	return &logCasoDeUso{
		contextTimeout: timeout,
		servicioLog:    servicioLog,
	}
}

func (uc logCasoDeUso) RegistroEvento(ctx context.Context, mensaje modelos.MensanjeEvento) error {
	return uc.servicioLog.RegistrarEvento(context.Background(), mensaje)
}
