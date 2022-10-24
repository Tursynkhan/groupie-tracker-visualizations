package delivery

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func (h *Handler) artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.ErrorHandler(w, r, errStatus{http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)})
		return
	}
	if r.URL.Path != "/open" {
		h.ErrorHandler(w, r, errStatus{http.StatusNotFound, http.StatusText(http.StatusNotFound)})
		return
	}
	ts, err := template.ParseFiles("./ui/html/artist.html")
	if err != nil {
		log.Println(err.Error())
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		h.ErrorHandler(w, r, errStatus{http.StatusNotFound, http.StatusText(http.StatusNotFound)})
		return
	}

	idUrl := strconv.Itoa(id)

	res1, err := h.service.IdArtist(idUrl)
	if err != nil {
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	res2, err := h.service.Relations(idUrl)
	if err != nil {
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
		return
	}
	res1.DatesLocation = res2.DatesLocations

	err = ts.Execute(w, res1)
	if err != nil {
		log.Println(err.Error())
		h.ErrorHandler(w, r, errStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
	}
}
