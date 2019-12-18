package camacho

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"controlAsistenciaAPI/servicios/camacho/client"
	"controlAsistenciaAPI/servicios/camacho/client/operations"
	"controlAsistenciaAPI/servicios/camacho/models"
	"time"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

type servicioEstadoUsuaarioCamacho struct {
	contextTimeout time.Duration
	apiClient      *client.CamachoAPIClient
}

func NuevoServicioEstadoUsuarioCamacho(timeout time.Duration, host string) controlasistencia.ServicioEstadoUsuario {
	transport := httptransport.New(host, "", nil)
	client := client.New(transport, strfmt.Default)
	return &servicioEstadoUsuaarioCamacho{
		contextTimeout: timeout,
		apiClient:      client,
	}
}

func (s servicioEstadoUsuaarioCamacho) TraerEstadoActualDelUsuario(ctx context.Context, usuario modelos.Usuario) (modelos.EstadoAsistencia, error) {
	params := operations.NewConsultarEstadoParams().WithIDUsuario(usuario.ID)
	estado, err := s.apiClient.Operations.ConsultarEstado(params)
	if err != nil {
		return modelos.EstadoAsistenciaNoRegistrado, err
	}
	switch estado.Payload.Estado {
	case models.EstadoUsuarioTrabajando:
		return modelos.EstadoAsistenciaTrabajando, nil
	case models.EstadoUsuarioDescanso:
		return modelos.EstadoAsistenciaDescanso, nil
	case models.EstadoUsuarioNoRegistrado:
		return modelos.EstadoAsistenciaNoRegistrado, nil
	}
	return modelos.EstadoAsistenciaNoRegistrado, nil
}
