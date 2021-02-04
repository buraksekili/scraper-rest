package data

import (
	"encoding/json"
	"io"
)

// ToJSON parses Go data structure to JSON.
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

// FromJSON converts JSON to Go data structure.
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
