package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello from the browser!"))
}
