package proxy

import (
	"strings"
	"testing"
	"time"
)

func TestProxy(t *testing.T) {
	proxy := NewProxyServer()
	
	// First request - should go to real server
	start := time.Now()
	response1 := proxy.HandleRequest("/api/users")
	duration1 := time.Since(start)
	
	expected := "Response from real server for: /api/users"
	if response1 != expected {
		t.Errorf("Expected %s, got %s", expected, response1)
	}
	
	// Should take some time (at least 100ms)
	if duration1 < 100*time.Millisecond {
		t.Error("First request should take at least 100ms")
	}
	
	// Second request - should be cached
	start = time.Now()
	response2 := proxy.HandleRequest("/api/users")
	duration2 := time.Since(start)
	
	expectedCached := "Cached response for: /api/users - Response from real server for: /api/users"
	if response2 != expectedCached {
		t.Errorf("Expected %s, got %s", expectedCached, response2)
	}
	
	// Should be much faster (cached)
	if duration2 > 10*time.Millisecond {
		t.Error("Cached request should be much faster")
	}
	
	// Different URL should go to real server again
	response3 := proxy.HandleRequest("/api/products")
	if !strings.Contains(response3, "Response from real server for: /api/products") {
		t.Error("Different URL should go to real server")
	}
}