package api

import (
	"context"
)

// HelloHandler implements the Handler interface for hello operations
type HelloHandler struct{}

// NewHelloHandler creates a new HelloHandler instance
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}

// PostHello implements postHello operation.
// Receives a hello message and returns it back.
// POST /hello
func (h *HelloHandler) PostHello(ctx context.Context, req *HelloRequest) (PostHelloRes, error) {
	// Simply echo back the received message
	response := &HelloResponse{
		Message: req.Message,
	}

	return response, nil
}
