package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/coder/websocket"
	"github.com/ryanwclark1/the-ui/cmd/web"
	"github.com/ryanwclark1/the-ui/cmd/web/pages/forms"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HelloWorldHandler)

	// Add the new route for the main page
  mux.HandleFunc("/main", s.mainHandler)

	mux.HandleFunc("/health", s.healthHandler)

	mux.HandleFunc("/websocket", s.websocketHandler)

	fileServer := http.FileServer(http.FS(web.Files))
	mux.Handle("/assets/", fileServer)
	mux.Handle("/web", templ.Handler(web.HelloForm()))
	mux.HandleFunc("/hello", web.HelloWebHandler)

	mux.HandleFunc("/forms/layout", s.LayoutFormHandler)
	mux.HandleFunc("/forms/input", s.InputHandler)
	mux.HandleFunc("/forms/checkbox", s.CheckBoxHandler)
	mux.HandleFunc("/forms/radio", s.RadioHandler)
	mux.HandleFunc("/forms/textarea", s.TextAreaHandler)
	mux.HandleFunc("/forms/switch", s.SwitchHandler)
	mux.HandleFunc("/forms/select", s.SelectHandler)

	return mux
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

// Handler for the LayoutForm component
func (s *Server) LayoutFormHandler(w http.ResponseWriter, r *http.Request) {
    component := LayoutForm()
    err := templ.Render(w, component)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// Handler for the Input component
func (s *Server) InputHandler(w http.ResponseWriter, r *http.Request) {
    component := InputForm()
    err := templ.Render(w, component)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

// func LayoutFormHandler(w http.ResponseWriter, r *http.Request) {
//     tmpl := template.Must(template.ParseFiles("templates/forms/layout.html"))
//     tmpl.Execute(w, nil)
// }


func (s *Server) mainHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	component := web.MainPage()
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in mainHandler: %e", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(s.db.Health())
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) websocketHandler(w http.ResponseWriter, r *http.Request) {
	socket, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Printf("could not open websocket: %v", err)
		_, _ = w.Write([]byte("could not open websocket"))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer socket.Close(websocket.StatusGoingAway, "server closing websocket")

	ctx := r.Context()
	socketCtx := socket.CloseRead(ctx)

	for {
		payload := fmt.Sprintf("server timestamp: %d", time.Now().UnixNano())
		err := socket.Write(socketCtx, websocket.MessageText, []byte(payload))
		if err != nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
}
