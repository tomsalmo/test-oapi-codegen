package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/tomsalmo/test-oapi-codegen/internal/api"
	"github.com/tomsalmo/test-oapi-codegen/internal/api/handlers"
)

func main() {
	ctx := context.Background()
	// Listen for SIGINT to gracefully shutdown.
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer stop()

	mux := http.NewServeMux()
	h := api.HandlerFromMux(api.NewStrictHandler(handlers.Handler{}, []nethttp.StrictHTTPMiddlewareFunc{}), mux)
	server := &http.Server{
		Addr: "localhost:9000",
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
		Handler: h,
	}
	if err := server.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
