package servicios

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"time"
)

type servicioRegistroDeEventosPumari struct {
	contextTimeout time.Duration
}

func NuevoServicioRegistroDeEventosPumari(timeout time.Duration) controlasistencia.ServicioRegistroEventos {
	return &servicioRegistroDeEventosPumari{
		contextTimeout: timeout,
	}
}

func (s servicioRegistroDeEventosPumari) RegistrarEntrada(ctx context.Context, usuario modelos.Usuario, date time.Time) error {
	return nil
}

func (s servicioRegistroDeEventosPumari) RegistrarSalida(ctx context.Context, usuario modelos.Usuario, date time.Time) error {
	return nil
}
