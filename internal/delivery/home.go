package delivery

import (
	"log"
	"net/http"
	"text/template"
)

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.ErrorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	if r.URL.Path != "/" {
		h.ErrorHandler(w, r, errStatus{http.StatusNotFound, http.StatusText(http.StatusNotFound)})
		return
	}
	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	res, err := h.service.Allartist()
	if err != nil {
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	err = ts.Execute(w, res)
	if err != nil {
		log.Println(err.Error())
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
	}
}
