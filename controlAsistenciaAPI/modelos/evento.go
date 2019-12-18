package modelos

import "time"

type Evento struct {
	ID            string
	NombreUsuario string
	Tipo          TipoEvento
	Fecha         time.Time
}
