package delivery

import (
	"main/internal/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.home)
	mux.HandleFunc("/open", h.artist)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
