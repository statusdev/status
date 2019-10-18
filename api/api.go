package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-openapi/loads"
	restmiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/statusdev/status/api/v1/models"
	"github.com/statusdev/status/api/v1/restapi"
	"github.com/statusdev/status/api/v1/restapi/operations"
	"github.com/statusdev/status/api/v1/restapi/operations/status"
	service "github.com/statusdev/status/status"
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
	var svc service.Service

	svc = service.NewService()
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

func NewSetStatusHandler(svc service.Service) status.SetStatusHandlerFunc {
	return func(params status.SetStatusParams) restmiddleware.Responder {

		body := params.Media

		// validate if media for status is valid

		// upload the media and receive media url

		// store new status item to backend

		s := models.Status{
			Caption: "testcaption",
			Media:   "https://foo.bar.com",
		}

		res, err := svc.AddStatus(s)
		if err != nil {
			return status.NewGetStatusInternalServerError()
		}
		return status.NewSetStatusOK().WithPayload(convertStatus(res))
	}
}

func NewGetStatusHandler(svc service.Service) status.GetStatusHandlerFunc {
	return func(params status.GetStatusParams) restmiddleware.Responder {
		res, err := svc.GetStatus()
		if err != nil {
			return status.NewGetStatusInternalServerError()
		}
		return status.NewSetStatusOK().WithPayload(convertStatusList(res))
	}
}


func NewNotifyHandler(svc service.Service) status.NotifyHandlerFunc {
	return func(params status.NotifyParams) restmiddleware.Responder {
		body := params.Body

		svc.UpdateSubscriptionFrom()
		if err != nil {
			return status.NewGetStatusInternalServerError()
		}
		return status.NewSetStatusOK().WithPayload(convertStatusList(res))
	}
}
