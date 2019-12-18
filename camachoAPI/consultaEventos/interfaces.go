package consultaEventos

import (
	"camachoAPI/modelos"
	"context"
)

type CasoDeUso interface {
	ConsultarEventos(ctx context.Context, usuarioID string) ([]modelos.Evento, error)
	ConsultarEstado(ctx context.Context, usuarioID string) (modelos.EstadoAsistencia, error)
}

type RepositorioEventos interface {
	ListaDeEventos(ctx context.Context, usuarioID string) ([]modelos.Evento, error)
	EstadoUsuario(ctx context.Context, usuarioID string) (modelos.EstadoAsistencia, error)
}
