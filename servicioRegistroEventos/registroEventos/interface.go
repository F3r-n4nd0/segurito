package registroEventos

import "context"

type CasoDeUso interface {
	RegistrarEvento(ctx context.Context, codigoUsuario string) error
}

type ServicioLectorDeMensajes interface {
	LeerMensajeDeEventos(ctx context.Context, codigoUsuario string)  error
}