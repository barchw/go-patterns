package chain

import (
	"fmt"
	"time"
)

type Request struct {
	Headers map[string]string
	Path    string
	Body    string
}

type Response struct {
	StatusCode int
	Body       string
}

type Handler interface {
	Handle(request *Request) *Response
	SetNext(Handler)
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) SetNext(next Handler) {
	h.next = next
}

func (h *BaseHandler) handleNext(request *Request) *Response {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return &Response{StatusCode: 200, Body: "OK"}
}

type AuthHandler struct {
	BaseHandler
	token string
}

func NewAuthHandler(token string) *AuthHandler {
	return &AuthHandler{token: token}
}

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

type LoggingHandler struct {
	BaseHandler
	Logs []string
}

func NewLoggingHandler() *LoggingHandler {
	return &LoggingHandler{}
}

func (h *LoggingHandler) Handle(request *Request) *Response {
	h.Logs = append(h.Logs, fmt.Sprintf("%s %s", time.Now().Format(time.RFC3339), request.Path))
	return h.handleNext(request)
}

type RateLimitHandler struct {
	BaseHandler
	maxRequests int
	count       int
}

func NewRateLimitHandler(maxRequests int) *RateLimitHandler {
	return &RateLimitHandler{maxRequests: maxRequests}
}

func (h *RateLimitHandler) Handle(request *Request) *Response {
	h.count++
	if h.count > h.maxRequests {
		return &Response{StatusCode: 429, Body: "Too Many Requests"}
	}
	return h.handleNext(request)
}

func BuildChain(handlers ...Handler) Handler {
	for i := 0; i < len(handlers)-1; i++ {
		handlers[i].SetNext(handlers[i+1])
	}
	if len(handlers) > 0 {
		return handlers[0]
	}
	return nil
}
