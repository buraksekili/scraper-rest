package handlers

import (
	"log"
	"net/http"
)

type Info struct {
	l *log.Logger
}

func GetNewInfo(l *log.Logger) *Info{
	return &Info{l: l}
}

func (i *Info) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}