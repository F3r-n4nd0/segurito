package mesa

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"os/exec"
	"strings"
	"time"
)

type servicioAutentificacionMesa struct {
	contextTimeout    time.Duration
	path              string
	nombreApplicacion string
}

func NuevoServicioDeAutentificacionMesa(timeout time.Duration, path string) controlasistencia.ServicioAutentification {
	return &servicioAutentificacionMesa{
		contextTimeout:    timeout,
		path:              path,
		nombreApplicacion: "mesa",
	}
}

func (s servicioAutentificacionMesa) AutentificarUsuario(ctx context.Context, codigoUsuario string) (*modelos.Usuario, error) {
	strPath := s.path + s.nombreApplicacion
	cmd := exec.Command(strPath, codigoUsuario)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	response := string(out)
	if response == "Unknown" {
		return nil, &modelos.ErrorCodigoUsuarioInvalido{}
	}
	strLines := strings.Split(response, "\n")
	user := modelos.Usuario{
		ID:   strings.Replace(strLines[0], "ID: ", "", -1),
		Name: strings.Replace(strLines[1], "User: ", "", -1),
	}
	return &user, nil
}
