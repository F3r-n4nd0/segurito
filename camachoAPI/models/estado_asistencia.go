package models

type EstadoAsistencia string

const (
	EstadoAsistenciaTrabajando   EstadoAsistencia = "Trabajando"
	EstadoAsistenciaDescanso                      = "Descanso"
	EstadoAsistenciaNoRegistrado                  = "No registrado"
)
