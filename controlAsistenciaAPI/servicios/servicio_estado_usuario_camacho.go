package servicios

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"time"
)

type servicioEstadoUsuaarioCamacho struct {
	contextTimeout time.Duration
}

func NuevoServicioEstadoUsuarioCamacho(timeout time.Duration) controlasistencia.ServicioEstadoUsuario {
	return &servicioEstadoUsuaarioCamacho{
		contextTimeout: timeout,
	}
}

func (s servicioEstadoUsuaarioCamacho) TraerEstadoActualDelUsuario(ctx context.Context, usuario modelos.Usuario) (modelos.EstadoAsistencia, error) {
	return modelos.EstadoAsistenciaNoRegistrado, nil
}
