package consultaEventos

import (
	"camachoAPI/modelos"
	"context"
	"time"
)

type consultaEventosCasoDeUso struct {
	contextTimeout     time.Duration
	repositorioEventos RepositorioEventos
}

func NuevoCasoDeUsoConsultaEventos(timeout time.Duration,
	repositorioEventos RepositorioEventos,
) CasoDeUso {
	return &consultaEventosCasoDeUso{
		contextTimeout:     timeout,
		repositorioEventos: repositorioEventos,
	}
}

func (c consultaEventosCasoDeUso) ConsultarEventos(ctx context.Context, usuarioID string) ([]modelos.Evento, error) {
	listaEventos, err := c.repositorioEventos.ListaDeEventos(ctx, usuarioID)
	if err != nil {
		return nil, err
	}
	return listaEventos, nil
}

func (c consultaEventosCasoDeUso) ConsultarEstado(ctx context.Context, usuarioID string) (modelos.EstadoAsistencia, error) {
	estado, err := c.repositorioEventos.EstadoUsuario(ctx, usuarioID)
	if err != nil {
		return modelos.EstadoAsistenciaNoRegistrado, err
	}
	return estado, nil
}
