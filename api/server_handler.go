package api

import "net/http"

type ServerHandler struct {
}

func (h *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Handler(w, r)
}
