package handlers

import "fmt"

var ErrInvalidURL = fmt.Errorf("Given URL is unavailable for now.")

type HandlerError struct {
	Message string `json:"message"`
}
