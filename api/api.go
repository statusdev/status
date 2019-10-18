package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-openapi/loads"
	restmiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/statusdev/status/api/v1/restapi"
	"github.com/statusdev/status/api/v1/restapi/operations"
	"net/http"
)

func NewStatusAPI(logger log.Logger) (*chi.Mux, error) {

	router := chi.NewRouter()

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, fmt.Errorf("failed to load embedded swagger file: %w", err.Error())
	}

	api := operations.NewStatusapiAPI(swaggerSpec)

	// Skip the  redoc middleware, only serving the OpenAPI specification and
	// the API itself via RoutesHandler. See:
	// https://github.com/go-swagger/go-swagger/issues/1779
	api.Middleware = func(b restmiddleware.Builder) http.Handler {
		return restmiddleware.Spec("", swaggerSpec.Raw(), api.Context().RoutesHandler(b))
	}

	// initialize services
	//var svc service.Service

	//svc = service.NewService()
	//if err != nil {
	//	return nil, err
	//}

	//svc = service.NewLoggingService(log.WithPrefix(logger, "service", "svc"), svc)

	// namespaces
	api.StatusGetStatusHandler = nil
	api.StatusSetStatusHandler = nil
	api.StatusNotifyHandler = nil

	api.SubscribeSubscribeHandler = nil
	api.SubscribeUnsubscribeHandler = nil

	router.Mount("/", api.Serve(nil))

	return router, nil
}