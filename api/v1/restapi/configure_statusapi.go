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
	"github.com/statusdev/status/api/v1/restapi/operations/subscribe"
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
	if api.StatusSetStatusHandler == nil {
		api.StatusSetStatusHandler = status.SetStatusHandlerFunc(func(params status.SetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation status.SetStatus has not yet been implemented")
		})
	}
	if api.SubscribeSubscribeHandler == nil {
		api.SubscribeSubscribeHandler = subscribe.SubscribeHandlerFunc(func(params subscribe.SubscribeParams) middleware.Responder {
			return middleware.NotImplemented("operation subscribe.Subscribe has not yet been implemented")
		})
	}
	if api.SubscribeUnsubscribeHandler == nil {
		api.SubscribeUnsubscribeHandler = subscribe.UnsubscribeHandlerFunc(func(params subscribe.UnsubscribeParams) middleware.Responder {
			return middleware.NotImplemented("operation subscribe.Unsubscribe has not yet been implemented")
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
