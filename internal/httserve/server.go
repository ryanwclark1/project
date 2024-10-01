package httpserver

import (
	"net/http"
)

// NewServer constructs the main HTTP handler with all middleware applied.
func NewServer(
	logger *Logger
	config *Config
	tenantsStore *tenantsStore
	commentsStore *commentsStore
	conversationService *conversationService
	chatGPTService *chatGPTService
	authProxy *authProxy
) http.Handler {
    mux := http.NewServeMux()

    // Add routes
    addRoutes(
			mux,
			Logger,
			Config,
			tenantsStore,
			commentsStore,
			conversationService,
			chatGPTService,
			authProxy,
		)

    // Apply middleware
    var handler http.Handler = mux
		handler = loggingMiddleware(handler)
		handler = authMiddleware(handler)
    return handler
}

// Example middleware (logging)
func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Log request details
        println("Request: ", r.Method, r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

// Example middleware (auth)
func authMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Authenticate the request
        next.ServeHTTP(w, r)
    })
}

// tenantsStore middleware
func tenantsStoreMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Authenticate the request
				next.ServeHTTP(w, r)
		})
}

// commentsStore middleware
func commentsStoreMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Authenticate the request
				next.ServeHTTP(w, r)
		})
}

// conversationService middleware
func conversationServiceMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Authenticate the request
				next.ServeHTTP(w, r)
		})
}

// chatGPTService middleware
func chatGPTServiceMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Authenticate the request
				next.ServeHTTP(w, r)
		})
}

// authProxy middleware
func authProxyMiddleware(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Authenticate the request
				next.ServeHTTP(w, r)
		})
}