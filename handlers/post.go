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
	switch err {
	case nil:
	case ErrInvalidURL:
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(&HandlerError{Message: ErrInvalidURL.Error()}, w)
		return
	default:
		w.WriteHeader(http.StatusInternalServerError)
		err2 := data.ToJSON(&HandlerError{Message: err.Error()}, w)
		if err2 != nil {
			i.l.Println("err2 as ", err2)
		}
		return
	}

	err = data.ToJSON(images, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&HandlerError{Message: "Serialization error! Try again."}, w)
		return
	}
	return
}
