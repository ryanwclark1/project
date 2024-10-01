package httpserver

import (
	"net/http"
	"log/slog"
)

// addRoutes maps all the API endpoints and health check routes
func addRoutes(
	mux                    *http.ServeMux,
	logger                 *loggin.Logger,
	config                 Config,
	tenantsStore           *TenantsStore,
	commentsStore          *CommentsStore,
	conversationService    *ConversationService,
	chatGPTService         *ChatGPTService,
	authProxy              *authProxy
) {
    mux.HandleFunc("/healthz", handleHealthzPlease(logger))
    mux.Handle("/api/v1/example", handleExample())  // Example route
		mux.Handle("/", http.NotFoundHandler())
}

// handleHealthzPlease returns an http.Handler for the health check endpoint with logging
func handleHealthzPlease(logger *slog.Logger) http.Handler {
    return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
        logger.Info("Health check requested", "method", r.Method, "path", r.URL.Path)
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Healthy"))
    })
}

// handleExample returns an http.Handler for an example API endpoint
func handleExample() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Example route"))
    })
}
