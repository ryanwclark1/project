package web

import (
	"log"
	"net/http"
)


func MainWebHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
    }

    component := MainPage()
    err = component.Render(r.Context(), w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        log.Fatalf("Error rendering in MainWebHandler: %e", err)
    }
}