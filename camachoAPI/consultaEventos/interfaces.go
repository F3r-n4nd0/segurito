package consultaEventos

import (
	"camachoAPI/models"
	"context"
)

type CasoDeUso interface {
	ConsultarEventos(ctx context.Context, usuarioID string) ([]models.Evento, error)
	ConsultarEstado(ctx context.Context, usuarioID string) (models.EstadoAsistencia, error)
}

type RepositorioEventos interface {
	ListaDeEventos(ctx context.Context, usuarioID string) ([]models.Evento, error)
	EstadoUsuario(ctx context.Context, usuarioID string) (models.EstadoAsistencia, error)
}
