package httpserver

import (
    "embed"
    "net/http"
    "log/slog"


    "github.com/coder/websocket"
    "github.com/a-h/templ"

    "github.com/ryanwclark1/the-ui/web"
    "github.com/ryanwclark1/the-ui/templates"
    "github.com/ryanwclark1/the-ui/pkgs/htmx"
)

import (
	"net/http"
	"log/slog"
    "io/fs"

	"github.com/coder/websocket"
    "github.com/a-h/templ"

	"github.com/ryanwclark1/the-ui/web"
    "github.com/ryanwclark1/the-ui/templates"
    "github.com/ryanwclark1/the-ui/pkgs/htmx"
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
		// Serve static files from ./web for CSS, JS, and images
    // fileServer := http.FileServer(http.Dir(web.Files))
    // mux.Handle("/static/", http.StripPrefix("/static", fileServer))
    // Static file handler for serving embedded files from the `web/` directory
    mux.Handle("/static/", http.StripPrefix("/static/", handleStaticFile()))

    mux.HandleFunc("/healthz", handleHealthzPlease(logger))
    mux.Handle("/api/v1/example", handleExample())  // Example route
	mux.Handle("/", http.NotFoundHandler())
    mux.Handle("/main", handleMain(logger))
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

// handleWebSocket returns an http.Handler that manages WebSocket connections.
func handleWebSocket(logger *slog.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Accept the WebSocket connection
        socket, err := websocket.Accept(w, r, nil)
        if err != nil {
            logger.Error("Could not open WebSocket", "error", err)
            http.Error(w, "could not open websocket", http.StatusInternalServerError)
            return
        }

        // Close the WebSocket when done
        defer func() {
            logger.Info("Closing WebSocket connection", "remote_addr", r.RemoteAddr)
            socket.Close(websocket.StatusGoingAway, "server closing websocket")
        }()

        logger.Info("WebSocket connection established", "remote_addr", r.RemoteAddr)

        // WebSocket loop for sending server timestamps
        ctx := r.Context()
        socketCtx := socket.CloseRead(ctx)

        for {
            payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
            err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
            if err != nil {
                logger.Error("Error sending WebSocket message", "error", err)
                break
            }

            logger.Info("WebSocket message sent", "payload", payload)
            time.Sleep(2 * time.Second)
        }

        logger.Info("Closing WebSocket connection", "remote_addr", r.RemoteAddr)
    })
}

func handleMain(logger *slog.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if htmx.IsHTMX(r) {
            // Handle HTMX request
            htmx.NewResponse().
                Reswap(htmx.SwapBeforeEnd).
                Retarget("#contacts").
                AddTrigger(htmx.Trigger("enable-submit")).
                AddTrigger(htmx.TriggerDetail("display-message", "Hello, HTMX!")).
                Write(w)
        } else {
            // Handle non-HTMX request by rendering the main template
            templates.RenderMainPage(r.Context(), w, "This is the main page rendered with Templ.")
        }
    })
}

// StaticFileHandler returns an http.Handler that serves the embedded static files.
func handleStaticFile() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Create a sub-filesystem by stripping the leading "web" directory from the paths
        fsys, err := fs.Sub(staticFiles, "web")
        if err != nil {
            http.Error(w, "could not load static files", http.StatusInternalServerError)
            return
        }

        // Serve static files using http.FileServer
        http.FileServer(http.FS(fsys)).ServeHTTP(w, r)
    })
}