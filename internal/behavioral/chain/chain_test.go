package chain

import "testing"

func TestChain_FullPassThrough(t *testing.T) {
	auth := NewAuthHandler("secret")
	logging := NewLoggingHandler()
	rateLimit := NewRateLimitHandler(10)

	chain := BuildChain(auth, logging, rateLimit)

	req := &Request{
		Headers: map[string]string{"Authorization": "Bearer secret"},
		Path:    "/api/data",
		Body:    "request body",
	}

	resp := chain.Handle(req)
	if resp.StatusCode != 200 {
		t.Errorf("StatusCode = %d, want 200", resp.StatusCode)
	}
	if len(logging.Logs) != 1 {
		t.Errorf("expected 1 log entry, got %d", len(logging.Logs))
	}
}

func TestChain_AuthShortCircuits_MissingToken(t *testing.T) {
	auth := NewAuthHandler("secret")
	logging := NewLoggingHandler()

	chain := BuildChain(auth, logging)

	req := &Request{
		Headers: map[string]string{},
		Path:    "/api/data",
	}

	resp := chain.Handle(req)
	if resp.StatusCode != 401 {
		t.Errorf("StatusCode = %d, want 401", resp.StatusCode)
	}
	if len(logging.Logs) != 0 {
		t.Errorf("logging should not have been reached, got %d entries", len(logging.Logs))
	}
}

func TestChain_AuthShortCircuits_InvalidToken(t *testing.T) {
	auth := NewAuthHandler("secret")
	logging := NewLoggingHandler()

	chain := BuildChain(auth, logging)

	req := &Request{
		Headers: map[string]string{"Authorization": "Bearer wrong"},
		Path:    "/api/data",
	}

	resp := chain.Handle(req)
	if resp.StatusCode != 403 {
		t.Errorf("StatusCode = %d, want 403", resp.StatusCode)
	}
	if len(logging.Logs) != 0 {
		t.Errorf("logging should not have been reached, got %d entries", len(logging.Logs))
	}
}

func TestChain_RateLimitShortCircuits(t *testing.T) {
	rateLimit := NewRateLimitHandler(2)
	chain := BuildChain(rateLimit)

	req := &Request{
		Headers: map[string]string{},
		Path:    "/api/data",
	}

	resp := chain.Handle(req)
	if resp.StatusCode != 200 {
		t.Errorf("first request: StatusCode = %d, want 200", resp.StatusCode)
	}

	resp = chain.Handle(req)
	if resp.StatusCode != 200 {
		t.Errorf("second request: StatusCode = %d, want 200", resp.StatusCode)
	}

	resp = chain.Handle(req)
	if resp.StatusCode != 429 {
		t.Errorf("third request: StatusCode = %d, want 429", resp.StatusCode)
	}
}

func TestChain_LoggingRecordsAllRequests(t *testing.T) {
	logging := NewLoggingHandler()
	chain := BuildChain(logging)

	paths := []string{"/a", "/b", "/c"}
	for _, p := range paths {
		chain.Handle(&Request{Headers: map[string]string{}, Path: p})
	}

	if len(logging.Logs) != 3 {
		t.Errorf("expected 3 log entries, got %d", len(logging.Logs))
	}
}
