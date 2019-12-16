package modelos

type ErrorEstadoUsuarioInvalido struct{}

func (m *ErrorEstadoUsuarioInvalido) Error() string {
	return "usuario en estado incorrecto para registro "
}
