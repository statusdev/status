// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/statusdev/status/api/v1/restapi/operations"
	"github.com/statusdev/status/api/v1/restapi/operations/status"
	"github.com/statusdev/status/api/v1/restapi/operations/subscribers"
	"github.com/statusdev/status/api/v1/restapi/operations/subscribtions"
)

//go:generate swagger generate server --target ../../v1 --name Statusapi --spec ../../../swagger.yaml --exclude-main

func configureFlags(api *operations.StatusapiAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.StatusapiAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.MultipartformConsumer = runtime.DiscardConsumer

	api.JSONProducer = runtime.JSONProducer()

	if api.StatusAddStatusHandler == nil {
		api.StatusAddStatusHandler = status.AddStatusHandlerFunc(func(params status.AddStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation status.AddStatus has not yet been implemented")
		})
	}
	if api.SubscribersAddSubscriberHandler == nil {
		api.SubscribersAddSubscriberHandler = subscribers.AddSubscriberHandlerFunc(func(params subscribers.AddSubscriberParams) middleware.Responder {
			return middleware.NotImplemented("operation subscribers.AddSubscriber has not yet been implemented")
		})
	}
	if api.SubscribtionsAddSubscriptionHandler == nil {
		api.SubscribtionsAddSubscriptionHandler = subscribtions.AddSubscriptionHandlerFunc(func(params subscribtions.AddSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation subscribtions.AddSubscription has not yet been implemented")
		})
	}
	if api.StatusGetStatusHandler == nil {
		api.StatusGetStatusHandler = status.GetStatusHandlerFunc(func(params status.GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation status.GetStatus has not yet been implemented")
		})
	}
	if api.StatusNotifyHandler == nil {
		api.StatusNotifyHandler = status.NotifyHandlerFunc(func(params status.NotifyParams) middleware.Responder {
			return middleware.NotImplemented("operation status.Notify has not yet been implemented")
		})
	}
	if api.SubscribersRemoveSubscriberHandler == nil {
		api.SubscribersRemoveSubscriberHandler = subscribers.RemoveSubscriberHandlerFunc(func(params subscribers.RemoveSubscriberParams) middleware.Responder {
			return middleware.NotImplemented("operation subscribers.RemoveSubscriber has not yet been implemented")
		})
	}
	if api.SubscribtionsRemoveSubscriptionHandler == nil {
		api.SubscribtionsRemoveSubscriptionHandler = subscribtions.RemoveSubscriptionHandlerFunc(func(params subscribtions.RemoveSubscriptionParams) middleware.Responder {
			return middleware.NotImplemented("operation subscribtions.RemoveSubscription has not yet been implemented")
		})
	}

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
