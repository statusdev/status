package api

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-kit/kit/log"
	"github.com/go-openapi/loads"
	restmiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/statusdev/status/api/v1/models"
	"github.com/statusdev/status/api/v1/restapi"
	"github.com/statusdev/status/api/v1/restapi/operations"
	"github.com/statusdev/status/api/v1/restapi/operations/status"
	"github.com/statusdev/status/api/v1/restapi/operations/subscribers"
	"github.com/statusdev/status/api/v1/restapi/operations/subscriptions"
	service "github.com/statusdev/status/status"
	"net/http"
)

func NewStatusAPI(publicAddr string, alias string, logger log.Logger) (*chi.Mux, error) {

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

	svc = service.NewService(publicAddr, alias)
	//svc = service.NewLoggingService(log.WithPrefix(logger, "service", "svc"), svc)

	// namespaces
	api.StatusGetStatusHandler = NewGetStatusHandler(svc)
	api.StatusAddStatusHandler = NewAddStatusHandler(svc)
	api.StatusNotifyHandler = NewNotifyHandler(svc)

	api.SubscribersAddSubscriberHandler = NewAddSubscriberHandler(svc)
	api.SubscribersRemoveSubscriberHandler = NewRemoveSubscriberHandler(svc)

	api.SubscriptionsAddSubscriptionHandler = NewAddSubscriptionHandler(svc)
	api.SubscriptionsRemoveSubscriptionHandler = NewRemoveSubscriptionHandler(svc)

	router.Mount("/", api.Serve(nil))

	return router, nil
}

func NewAddStatusHandler(svc service.Service) status.AddStatusHandlerFunc {
	return func(params status.AddStatusParams) restmiddleware.Responder {

		// body := params.Media

		// validate if media for status is valid

		// upload the media and receive media url

		// store new status item to backend

		err := svc.AddStatus(service.StatusItem{
			Media:   "placeholder",
			Caption: "foo bar",
		})

		if err != nil {
			fmt.Print("AAAAAAAAAAAAAAAAAAAAAAa")
			fmt.Printf("%s", err)
			return status.NewAddStatusInternalServerError()
		}
		return status.NewAddStatusOK()
	}
}

func NewGetStatusHandler(svc service.Service) status.GetStatusHandlerFunc {
	return func(params status.GetStatusParams) restmiddleware.Responder {
		res, err := svc.GetStatus()
		if err != nil {
			return status.NewGetStatusInternalServerError()
		}
		return status.NewGetStatusOK().WithPayload(ProfileStatusToModelList(res))
	}
}

func NewNotifyHandler(svc service.Service) status.NotifyHandlerFunc {
	return func(params status.NotifyParams) restmiddleware.Responder {
		body := params.Body

		err := svc.UpdateSubscriptionFrom(service.ProfileStatus{
			URL:    body.URL,
			Alias:  body.Alias,
			Status: ModeltoStatusItemList(body.Status),
		})
		if err != nil {
			return status.NewNotifyInternalServerError()
		}
		return status.NewNotifyOK()
	}
}

func NewAddSubscriberHandler(svc service.Service) subscribers.AddSubscriberHandlerFunc {
	return func(params subscribers.AddSubscriberParams) restmiddleware.Responder {
		body := params.Body

		err := svc.AddSubscriber(service.Profile{URL: body.URL})
		if err != nil {
			return subscribers.NewAddSubscriberInternalServerError()
		}
		return subscribers.NewAddSubscriberOK()
	}
}

func NewRemoveSubscriberHandler(svc service.Service) subscribers.RemoveSubscriberHandlerFunc {
	return func(params subscribers.RemoveSubscriberParams) restmiddleware.Responder {
		body := params.Body

		err := svc.RemoveSubscriber(service.Profile{URL: body.URL})
		if err != nil {
			return subscribers.NewRemoveSubscriberInternalServerError()
		}
		return subscribers.NewRemoveSubscriberOK()
	}
}

func NewAddSubscriptionHandler(svc service.Service) subscriptions.AddSubscriptionHandlerFunc {
	return func(params subscriptions.AddSubscriptionParams) restmiddleware.Responder {
		body := params.Body

		err := svc.SubscribeTo(service.Profile{URL: body.URL})
		if err != nil {
			return subscriptions.NewAddSubscriptionInternalServerError()
		}
		return subscriptions.NewAddSubscriptionOK()
	}
}

func NewRemoveSubscriptionHandler(svc service.Service) subscriptions.RemoveSubscriptionHandlerFunc {
	return func(params subscriptions.RemoveSubscriptionParams) restmiddleware.Responder {
		body := params.Body

		err := svc.UnsubscribeFrom(service.Profile{URL: body.URL})
		if err != nil {
			return subscriptions.NewRemoveSubscriptionInternalServerError()
		}
		return subscribers.NewRemoveSubscriberOK()
	}
}

func ProfileStatusToModelList(items []*service.ProfileStatus) []*models.ProfileStatus {
	var result []*models.ProfileStatus
	for _, item := range items {
		result = append(result, &models.ProfileStatus{
			URL:    item.URL,
			Alias:  item.Alias,
			Status: StatusItemToModelList(item.Status),
		})
	}
	return result
}

func StatusItemToModelList(items []service.StatusItem) []*models.StatusItem {
	var result []*models.StatusItem
	for _, item := range items {
		result = append(result, &models.StatusItem{
			ID:      strfmt.UUID(item.ID),
			Media:   item.Media,
			Caption: item.Caption,
		})
	}
	return result
}

func ModeltoStatusItemList(items []*models.StatusItem) []service.StatusItem {
	var result []service.StatusItem
	for _, item := range items {
		result = append(result, service.StatusItem{
			ID:      string(item.ID),
			Media:   item.Media,
			Caption: item.Caption,
		})
	}
	return result
}
