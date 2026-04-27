/*
Package chain -- Chain of Responsibility

What is it?
Chain of Responsibility is a behavioral design pattern that lets you pass a request along a chain
of handlers. Each handler decides whether to process the request or pass it on to the next link
in the chain. The sender doesn't need to know which handler will ultimately handle its request.

When to use?
  - When a request can be handled by different handlers and you don't know upfront which one should handle it.
  - When you want to build processing pipelines (middleware), e.g., authorization -> logging -> rate limiting.
  - When you want to dynamically configure the order and composition of the chain at runtime.
  - When each handler should be able to stop processing (e.g., return an authorization error).

When NOT to use?
  - When every request must be handled by exactly one, predetermined handler.
  - When the processing order doesn't matter -- a simple list of procedures is better.
  - When you require a guarantee of handling -- in a chain, there may be no handler to process the request.

Tips and pitfalls:
  - Chain vs Decorator: a Chain can stop processing; a Decorator always delegates further.
  - Embedding BaseHandler instead of inheritance -- the idiomatic approach in Go.
  - Handler order is crucial, e.g., AuthHandler should come before RateLimitHandler.
  - RateLimitHandler is stateful (count counter) -- in multithreaded systems it requires synchronization.
*/
package chain

import (
	"fmt"
	"time"
)

// Request represents an incoming HTTP request with headers, path, and body.
type Request struct {
	Headers map[string]string
	Path    string
	Body    string
}

// Response represents an HTTP response with a status code and body.
type Response struct {
	StatusCode int
	Body       string
}

// Handler defines the interface of a chain link with request handling and next-handler assignment.
type Handler interface {
	Handle(request *Request) *Response
	SetNext(Handler)
}

// BaseHandler provides shared logic for passing the request to the next handler in the chain.
type BaseHandler struct {
	next Handler
}

// SetNext sets the next link in the chain of responsibility.
func (h *BaseHandler) SetNext(next Handler) {
	h.next = next
}

func (h *BaseHandler) handleNext(request *Request) *Response {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return &Response{StatusCode: 200, Body: "OK"}
}

// AuthHandler checks the Authorization header and breaks the chain if the token is missing or invalid.
type AuthHandler struct {
	BaseHandler
	token string
}

// NewAuthHandler creates a new authorization handler with the expected token.
func NewAuthHandler(token string) *AuthHandler {
	return &AuthHandler{token: token}
}

// Handle verifies the authorization token. Passes the request on or returns a 401/403 error.
func (h *AuthHandler) Handle(request *Request) *Response {
	auth := request.Headers["Authorization"]
	if auth == "" {
		return &Response{StatusCode: 401, Body: "Unauthorized: missing token"}
	}
	if auth != "Bearer "+h.token {
		return &Response{StatusCode: 403, Body: "Forbidden: invalid token"}
	}
	return h.handleNext(request)
}

// LoggingHandler logs every request and always passes it on in the chain.
type LoggingHandler struct {
	BaseHandler
	// Logs stores the recorded log entries.
	Logs []string
}

// NewLoggingHandler creates a new logging handler.
func NewLoggingHandler() *LoggingHandler {
	return &LoggingHandler{}
}

// Handle logs the request path with a timestamp and passes it to the next link.
func (h *LoggingHandler) Handle(request *Request) *Response {
	h.Logs = append(h.Logs, fmt.Sprintf("%s %s", time.Now().Format(time.RFC3339), request.Path))
	return h.handleNext(request)
}

// RateLimitHandler limits the number of requests and rejects them after the limit is exceeded.
type RateLimitHandler struct {
	BaseHandler
	maxRequests int
	count       int
}

// NewRateLimitHandler creates a new rate limit handler with the specified maximum number of requests.
func NewRateLimitHandler(maxRequests int) *RateLimitHandler {
	return &RateLimitHandler{maxRequests: maxRequests}
}

// Handle checks the request limit. Passes through or returns a 429 error if the limit is exceeded.
func (h *RateLimitHandler) Handle(request *Request) *Response {
	h.count++
	if h.count > h.maxRequests {
		return &Response{StatusCode: 429, Body: "Too Many Requests"}
	}
	return h.handleNext(request)
}

// BuildChain links the given handlers into a chain of responsibility and returns the first link.
func BuildChain(handlers ...Handler) Handler {
	for i := 0; i < len(handlers)-1; i++ {
		handlers[i].SetNext(handlers[i+1])
	}
	if len(handlers) > 0 {
		return handlers[0]
	}
	return nil
}
