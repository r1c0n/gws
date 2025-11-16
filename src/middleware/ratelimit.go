package middleware

import (
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type RateLimitConfig struct {
	Enabled           bool
	RequestsPerMinute int
	Burst             int
	Whitelist         []string
	ExemptPaths       []string
	ErrorPagePath     string
}

type visitor struct {
	limiter  *rateLimiter
	lastSeen time.Time
}

type rateLimiter struct {
	tokens   int
	capacity int
	refill   int
	lastTime time.Time
	mu       sync.Mutex
}

var (
	visitors = make(map[string]*visitor)
	mu       sync.RWMutex
)

// newRateLimiter creates a token bucket rate limiter
func newRateLimiter(requestsPerMinute, burst int) *rateLimiter {
	return &rateLimiter{
		tokens:   burst,
		capacity: burst,
		refill:   requestsPerMinute,
		lastTime: time.Now(),
	}
}

// allow checks if a request is allowed based on token bucket algorithm
func (rl *rateLimiter) allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.lastTime)

	// Refill tokens based on time elapsed
	tokensToAdd := int(elapsed.Minutes() * float64(rl.refill))
	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.capacity {
			rl.tokens = rl.capacity
		}
		rl.lastTime = now
	}

	// Check if we have tokens available
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}

	return false
}

// getVisitor retrieves or creates a visitor for the given IP
func getVisitor(ip string, requestsPerMinute, burst int) *rateLimiter {
	mu.Lock()
	defer mu.Unlock()

	v, exists := visitors[ip]
	if !exists {
		limiter := newRateLimiter(requestsPerMinute, burst)
		visitors[ip] = &visitor{limiter: limiter, lastSeen: time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

// cleanupVisitors removes stale visitors (not seen in 3 minutes)
func cleanupVisitors() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, v := range visitors {
			if time.Since(v.lastSeen) > 3*time.Minute {
				delete(visitors, ip)
			}
		}
		mu.Unlock()
	}
}

// isWhitelisted checks if an IP is in the whitelist
func isWhitelisted(ip string, whitelist []string) bool {
	for _, whitelistedIP := range whitelist {
		if ip == whitelistedIP {
			return true
		}
	}
	return false
}

// getIP extracts the IP address from the request
func getIP(r *http.Request) string {
	// Check X-Forwarded-For header first (for proxies)
	forwarded := r.Header.Get("X-Forwarded-For")
	if forwarded != "" {
		// Take the first IP in the list
		return forwarded
	}

	// Fall back to RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return ip
}

func init() {
	go cleanupVisitors()
}

// RateLimitMiddleware limits the rate of requests per IP address
func RateLimitMiddleware(config RateLimitConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !config.Enabled {
				next.ServeHTTP(w, r)
				return
			}

			// Check if path is exempt (e.g., static assets)
			for _, exemptPath := range config.ExemptPaths {
				if strings.HasPrefix(r.URL.Path, exemptPath) {
					next.ServeHTTP(w, r)
					return
				}
			}

			ip := getIP(r)

			// Check if IP is whitelisted
			if isWhitelisted(ip, config.Whitelist) {
				next.ServeHTTP(w, r)
				return
			}

			// Get or create limiter for this IP
			limiter := getVisitor(ip, config.RequestsPerMinute, config.Burst)

			// Check if request is allowed
			if !limiter.allow() {
				// Try to serve custom 429 error page
				if config.ErrorPagePath != "" {
					if content, err := os.ReadFile(config.ErrorPagePath); err == nil {
						w.Header().Set("Content-Type", "text/html; charset=utf-8")
						w.WriteHeader(http.StatusTooManyRequests)
						w.Write(content)
						return
					}
				}
				// Fallback to plain text
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
