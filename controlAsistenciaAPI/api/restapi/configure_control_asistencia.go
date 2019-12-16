// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"controlAsistenciaAPI/controlasistencia"
	"controlAsistenciaAPI/modelos"
	"controlAsistenciaAPI/repositorio"
	"controlAsistenciaAPI/servicios"
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"controlAsistenciaAPI/api/restapi/operations"
)

//go:generate swagger generate server --target ../../api --name ControlAsistencia --spec ../swagger/swagger.yaml

func configureFlags(api *operations.ControlAsistenciaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ControlAsistenciaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	servicioAutentificacion := servicios.NuevoServicioDeAutentificacionMesa(2, "./mocks/")
	servicioRegistroEventos := servicios.NuevoServicioRegistroDeEventosPumari(2)
	servicioEstadoUsaurio := servicios.NuevoServicioEstadoUsuarioCamacho(2)
	repositorioEventos := repositorio.NuevoRepositorioEventosMongoDB()
	casoDeUsoControlAsistencia := controlasistencia.NuevoCasoDeUsoControlAsistencia(2,
		servicioAutentificacion,
		servicioRegistroEventos,
		servicioEstadoUsaurio,
		repositorioEventos)

	api.RegistrarEntradaHandler = operations.RegistrarEntradaHandlerFunc(func(params operations.RegistrarEntradaParams) middleware.Responder {
		err := casoDeUsoControlAsistencia.RegistroEntrada(context.Background(), *params.Usuario.CodigoUsuario)
		if err != nil {
			switch err.(type) {
			case *modelos.ErrorCodigoUsuarioInvalido:
				return operations.NewRegistrarEntradaUnauthorized()
			case *modelos.ErrorEstadoUsuarioInvalido:
				return operations.NewRegistrarEntradaConflict()
			default:
				return operations.NewRegistrarEntradaInternalServerError()
			}
		}
		return operations.NewRegsitrarEntradaCreated()
	})

	api.RegistrarSalidaHandler = operations.RegistrarSalidaHandlerFunc(func(params operations.RegistrarSalidaParams) middleware.Responder {
		err := casoDeUsoControlAsistencia.RegistroSalida(context.Background(), *params.Usuario.CodigoUsuario)
		if err != nil {
			switch err.(type) {
			case *modelos.ErrorCodigoUsuarioInvalido:
				return operations.NewRegistrarSalidaUnauthorized()
			case *modelos.ErrorEstadoUsuarioInvalido:
				return operations.NewRegistrarSalidaConflict()
			default:
				return operations.NewRegistrarSalidaInternalServerError()
			}
		}
		return operations.NewRegsitrarSalidaCreated()
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
	return handler
}
