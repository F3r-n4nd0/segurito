// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"camachoAPI/api/models"
	"camachoAPI/consultaEventos"
	"camachoAPI/modelos"
	"camachoAPI/repositorio"
	"context"
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/strfmt"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"camachoAPI/api/restapi/operations"
)

//go:generate swagger generate server --target ../../api --name Camacho --spec ../swagger/swagger.yaml

func configureFlags(api *operations.CamachoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CamachoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	mongoDBURL := getenv("MONGO_DB_URL", "mongodb://localhost:27017")

	repositoryEvents := repositorio.NuevoRepositorioEventosMongoDB(mongoDBURL)

	casoDeUseConsulta := consultaEventos.NuevoCasoDeUsoConsultaEventos(2, repositoryEvents)

	api.ConsultarEstadoHandler = operations.ConsultarEstadoHandlerFunc(func(params operations.ConsultarEstadoParams) middleware.Responder {
		estado, err := casoDeUseConsulta.ConsultarEstado(context.Background(), params.IDUsuario)
		if err != nil {
			log.Print(err.Error())
			return operations.NewConsultarEstadoBadRequest()
		}
		switch estado {
		case modelos.EstadoAsistenciaTrabajando:
			return operations.NewConsultarEstadoOK().WithPayload(&operations.ConsultarEstadoOKBody{
				Estado: models.EstadoUsuarioTrabajando,
			})
		case modelos.EstadoAsistenciaDescanso:
			return operations.NewConsultarEstadoOK().WithPayload(&operations.ConsultarEstadoOKBody{
				Estado: models.EstadoUsuarioDescanso,
			})
		case modelos.EstadoAsistenciaNoRegistrado:
			return operations.NewConsultarEstadoOK().WithPayload(&operations.ConsultarEstadoOKBody{
				Estado: models.EstadoUsuarioNoRegistrado,
			})
		}
		return operations.NewConsultarEstadoOK().WithPayload(&operations.ConsultarEstadoOKBody{
			Estado: models.EstadoUsuarioNoRegistrado,
		})
	})

	api.TraerEventosHandler = operations.TraerEventosHandlerFunc(func(params operations.TraerEventosParams) middleware.Responder {
		eventos, err := casoDeUseConsulta.ConsultarEventos(context.Background(), params.IDUsuario)
		if err != nil {
			log.Print(err.Error())
			return operations.NewTraerEventosBadRequest()
		}

		listEvents := make([]*models.Eventos, 0)

		for _, element := range eventos {
			eventDate := strfmt.DateTime(element.Fecha)
			eventId := strfmt.UUID(element.ID)
			var eventType models.TipoEvento
			switch element.Tipo {
			case modelos.EntradaTipoEvento:
				eventType = models.TipoEventoEntrada
			case modelos.SalidaTipoEvento:
				eventType = models.TipoEventoSalida
			}
			newEvent := models.Eventos{
				Fecha:   &eventDate,
				ID:      &eventId,
				Tipo:    eventType,
				Usuario: &element.NombreUsuario,
			}
			listEvents = append(listEvents, &newEvent)
		}

		return operations.NewTraerEventosOK().WithPayload(listEvents)

	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return uiMiddleware(handler)
}

func uiMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
