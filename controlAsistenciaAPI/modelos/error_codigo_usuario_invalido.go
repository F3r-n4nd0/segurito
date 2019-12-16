package modelos

type ErrorCodigoUsuarioInvalido struct{}

func (m *ErrorCodigoUsuarioInvalido) Error() string {
	return "codigo de usuario incorrecto"
}
