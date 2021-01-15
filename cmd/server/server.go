package main

import (
	"context"
	"fmt"
	"net/http"
)

type HTTPPortNumber int

// apiServer configures necessary handlers and starts listening on a configured port.
type apiServer struct {
	port    HTTPPortNumber
	handler HTTPHandlerFunc
	server  *http.Server
}

func (s *apiServer) start() error {
	if s.handler == nil {
		return fmt.Errorf("HTTP handler is not defined - cannot start")
	}
	if s.port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := new(http.ServeMux)
	handler.HandleFunc("/", s.handler)

	s.server = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.port),
		Handler: handler,
	}
	return s.server.ListenAndServe()
}

// stop will shut down previously started HTTP server.
func (s *apiServer) stop() error {
	if s.server == nil {
		return fmt.Errorf("server was not started")
	}

	return s.server.Shutdown(context.Background())
}
