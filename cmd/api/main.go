package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/oklog/run"
	"github.com/statusdev/status/api"
	"github.com/urfave/cli"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	FlagHTTPAddr = "http-addr"
	FlagLogLevel = "log-level"

	EnvHTTPAddr = "API_HTTP_ADDR"
	EnvLogLevel = "API_LOG_LEVEL"
)

type apiConf struct {
	HTTPAddr string
	LogLevel string
}

var (
	apiConfig = apiConf{}
	apiFlags  = []cli.Flag{
		cli.StringFlag{
			Name:        FlagHTTPAddr,
			EnvVar:      EnvHTTPAddr,
			Usage:       "The address status API runs on",
			Value:       ":6660",
			Destination: &apiConfig.HTTPAddr,
		},
		cli.StringFlag{
			Name:        FlagLogLevel,
			EnvVar:      EnvLogLevel,
			Usage:       "The log level to filter logs with before printing",
			Value:       "info",
			Destination: &apiConfig.LogLevel,
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "status-api"
	app.Usage = "public api for the status instance"
	app.Action = apiAction
	app.Flags = apiFlags

	if err := app.Run(os.Args); err != nil {
		logger := NewLogger(false, "info")
		level.Info(logger).Log("msg", "failed to run the api", "err", err)
		os.Exit(1)
	}
}

func apiAction(c *cli.Context) error {

	logger := NewLogger(false, apiConfig.LogLevel)
	logger = log.WithPrefix(logger, "app", c.App.Name)

	var gr run.Group
	{
		apiV1, err := api.NewStatusAPI(log.WithPrefix(logger, "component", "api"))
		if err != nil {
			return err
		}

		router := chi.NewRouter()
		router.Use(Logger(logger))
		router.Mount("/", apiV1)

		server := http.Server{
			Addr:    apiConfig.HTTPAddr,
			Handler: router,
		}

		gr.Add(func() error {
			level.Info(logger).Log(
				"msg", "running api",
				"addr", server.Addr,
			)
			return server.ListenAndServe()
		}, func(err error) {
			_ = server.Shutdown(context.TODO())
		})
	}

	if err := gr.Run(); err != nil {
		return fmt.Errorf("error running: %w", err)
	}

	return nil
}

// Logger proxies incoming requests and prints them to stdout
func Logger(logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			next.ServeHTTP(ww, r)

			// maybe add more params for tenants?
			level.Debug(logger).Log(
				"proto", r.Proto,
				"method", r.Method,
				"status", ww.Status(),
				"path", r.URL.Path,
				"duration", time.Since(start),
				"bytes", ww.BytesWritten(),
			)
		})
	}
}

// NewLogger returns a new logger for the given loglevel
func NewLogger(json bool, loglevel string) log.Logger {
	var logger log.Logger

	if json {
		logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	} else {
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	}

	switch strings.ToLower(loglevel) {
	case "debug":
		logger = level.NewFilter(logger, level.AllowDebug())
	case "warn":
		logger = level.NewFilter(logger, level.AllowWarn())
	case "error":
		logger = level.NewFilter(logger, level.AllowError())
	default:
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	return log.With(logger,
		"ts", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
}
