package main

import (
	"fmt"

	"claude-test/internal/behavioral/chain"
)

func main() {
	fmt.Println("=== Chain of Responsibility Pattern ===")

	auth := chain.NewAuthHandler("my-secret-token")
	logging := chain.NewLoggingHandler()
	rateLimit := chain.NewRateLimitHandler(3)

	handler := chain.BuildChain(auth, logging, rateLimit)

	fmt.Println("Request without auth:")
	resp := handler.Handle(&chain.Request{
		Headers: map[string]string{},
		Path:    "/api/users",
	})
	fmt.Printf("  Status: %d, Body: %s\n", resp.StatusCode, resp.Body)

	fmt.Println("Request with valid auth:")
	resp = handler.Handle(&chain.Request{
		Headers: map[string]string{"Authorization": "Bearer my-secret-token"},
		Path:    "/api/users",
	})
	fmt.Printf("  Status: %d, Body: %s\n", resp.StatusCode, resp.Body)

	fmt.Println("Request with invalid auth:")
	resp = handler.Handle(&chain.Request{
		Headers: map[string]string{"Authorization": "Bearer bad-token"},
		Path:    "/api/users",
	})
	fmt.Printf("  Status: %d, Body: %s\n", resp.StatusCode, resp.Body)
}
