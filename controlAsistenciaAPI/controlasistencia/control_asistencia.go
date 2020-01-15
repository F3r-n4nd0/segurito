package controlasistencia

import (
	"context"
	"controlAsistenciaAPI/modelos"
	"time"

	"github.com/google/uuid"
)

type controlAsistenciaCasoDeUso struct {
	contextTimeout          time.Duration
	servicioAutentification ServicioAutentification
	servicioRegistroEventos ServicioRegistroEventos
	servicioEstadoUsuario   ServicioEstadoUsuario
	repositorioEventos      RepositorioEventos
}

func NuevoCasoDeUsoControlAsistencia(timeout time.Duration,
	servicioAutentification ServicioAutentification,
	servicioRegistroEventos ServicioRegistroEventos,
	servicioEstadoUsuario ServicioEstadoUsuario,
	repositorioEventos RepositorioEventos,
) CasoDeUso {
	return &controlAsistenciaCasoDeUso{
		contextTimeout:          timeout,
		servicioAutentification: servicioAutentification,
		servicioRegistroEventos: servicioRegistroEventos,
		servicioEstadoUsuario:   servicioEstadoUsuario,
		repositorioEventos:      repositorioEventos,
	}
}

func (c controlAsistenciaCasoDeUso) RegistroEntrada(ctx context.Context, codigoUsuario string) error {
	usuario, err := c.servicioAutentification.AutentificarUsuario(ctx, codigoUsuario)
	if err != nil {
		return err
	}
	if usuario == nil {
		return &modelos.ErrorCodigoUsuarioInvalido{}
	}
	estado, err := c.servicioEstadoUsuario.TraerEstadoActualDelUsuario(ctx, *usuario)
	if err != nil {
		return err
	}
	switch estado {
	case modelos.EstadoAsistenciaDescanso,
		modelos.EstadoAsistenciaNoRegistrado:
		err := c.registrarEntrada(ctx, *usuario)
		if err != nil {
			return err
		}
	case modelos.EstadoAsistenciaTrabajando:
		return &modelos.ErrorEstadoUsuarioInvalido{}
	}
	return nil
}

func (c controlAsistenciaCasoDeUso) registrarEntrada(ctx context.Context, usuario modelos.Usuario) error {
	currentTime := time.Now()
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	evento := modelos.Evento{
		ID:            newUUID.String(),
		NombreUsuario: usuario.ID,
		Tipo:          modelos.EntradaTipoEvento,
		Fecha:         time.Now(),
	}
	err = c.repositorioEventos.AlmacenarEvento(ctx, evento)
	if err != nil {
		return err
	}
	err = c.servicioRegistroEventos.RegistrarEntrada(ctx, usuario, currentTime)
	if err != nil {
		return err
	}
	return nil
}

func (c controlAsistenciaCasoDeUso) RegistroSalida(ctx context.Context, codigoUsuario string) error {
	usuario, err := c.servicioAutentification.AutentificarUsuario(ctx, codigoUsuario)
	if err != nil {
		return err
	}
	if usuario == nil {
		return &modelos.ErrorCodigoUsuarioInvalido{}
	}
	estado, err := c.servicioEstadoUsuario.TraerEstadoActualDelUsuario(ctx, *usuario)
	if err != nil {
		return err
	}
	switch estado {
	case modelos.EstadoAsistenciaTrabajando,
		modelos.EstadoAsistenciaNoRegistrado:
		err := c.registrarSalida(ctx, *usuario)
		if err != nil {
			return err
		}
	case modelos.EstadoAsistenciaDescanso:
		return &modelos.ErrorEstadoUsuarioInvalido{}
	}
	return nil
}

func (c controlAsistenciaCasoDeUso) registrarSalida(ctx context.Context, usuario modelos.Usuario) error {
	currentTime := time.Now()
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	evento := modelos.Evento{
		ID:            newUUID.String(),
		NombreUsuario: usuario.ID,
		Tipo:          modelos.SalidaTipoEvento,
		Fecha:         time.Now(),
	}
	err = c.repositorioEventos.AlmacenarEvento(ctx, evento)
	if err != nil {
		return err
	}
	err = c.servicioRegistroEventos.RegistrarSalida(ctx, usuario, currentTime)
	if err != nil {
		return err
	}
	return nil
}
