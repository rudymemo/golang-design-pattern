package proxy

import (
	"fmt"
	"time"
)

type Server interface {
	HandleRequest(url string) string
}

type RealServer struct{}

func (rs *RealServer) HandleRequest(url string) string {
	// Simulate processing time
	time.Sleep(100 * time.Millisecond)
	return fmt.Sprintf("Response from real server for: %s", url)
}

type ProxyServer struct {
	realServer *RealServer
	cache      map[string]string
}

func NewProxyServer() *ProxyServer {
	return &ProxyServer{
		realServer: &RealServer{},
		cache:      make(map[string]string),
	}
}

func (ps *ProxyServer) HandleRequest(url string) string {
	// Check cache first
	if response, exists := ps.cache[url]; exists {
		return fmt.Sprintf("Cached response for: %s - %s", url, response)
	}
	
	// If not in cache, forward to real server
	response := ps.realServer.HandleRequest(url)
	
	// Cache the response
	ps.cache[url] = response
	
	return response
}

func (ps *ProxyServer) ClearCache() {
	ps.cache = make(map[string]string)
}