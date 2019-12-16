package controlasistencia

import (
	"context"
	"controlAsistenciaAPI/modelos"
	"time"
)

type CasoDeUso interface {
	RegistroEntrada(ctx context.Context, codigoUsuario string) error
	RegistroSalida(ctx context.Context, codigoUsuario string) error
}

type ServicioAutentification interface {
	AutentificarUsuario(ctx context.Context, codigoUsuario string) (*modelos.Usuario, error)
}

type ServicioRegistroEventos interface {
	RegistrarEntrada(ctx context.Context, usuario modelos.Usuario, date time.Time) error
	RegistrarSalida(ctx context.Context, usuario modelos.Usuario, date time.Time) error
}

type RepositorioEventos interface {
	AlmacenarEntrada(ctx context.Context, usuario modelos.Usuario, date time.Time) error
	AlmacenarSalida(ctx context.Context, usuario modelos.Usuario, date time.Time) error
}

type ServicioEstadoUsuario interface {
	TraerEstadoActualDelUsuario(ctx context.Context, usuario modelos.Usuario) (modelos.EstadoAsistencia, error)
}
