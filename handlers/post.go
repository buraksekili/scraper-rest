package handlers

import (
	"net/http"

	"github.com/buraksekili/scraper-rest/data"
)

func (i *Info) ParseImages(w http.ResponseWriter, r *http.Request) {
	url := getURL(w, r)
	if url == "" || len(url) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&HandlerError{Message: ErrInvalidURL.Error()}, w)
		return
	}

	images, err := fetchURL(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&HandlerError{Message: err.Error()}, w)
		return
	}
	switch err {
	case nil:
	case ErrInvalidURL:
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&HandlerError{Message: ErrInvalidURL.Error()}, w)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&HandlerError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(images, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&HandlerError{Message: "Serialization error! Try again."}, w)
		return
	}
}
