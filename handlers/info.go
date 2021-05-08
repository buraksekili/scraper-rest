package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/buraksekili/scraper-rest/data"
	"golang.org/x/net/html"
)

// Info holds HTTP handler for getting page information.
type Info struct {
	l *log.Logger
}

// GetNewInfo creates handler with logger for HTTP requests
// Constructor for Info
func GetNewInfo(l *log.Logger) *Info {
	return &Info{l: l}
}

// fetchURL fetches the images of given URL
func fetchURL(URL string) (data.Images, error) {

	resp, err := http.Head(URL)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode > 399 {
		return nil, fmt.Errorf("cannot issue HEAD to %v: response status code %d",
			URL, resp.StatusCode)
	}

	resp, err = http.DefaultClient.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, ErrInvalidURL
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing html error: %s: %s\n", URL, resp.Status)
	}

	return data.GetImages(nil, doc, resp), nil
}

// getURL extracts URL from JSON taken from r.Body
func getURL(w http.ResponseWriter, r *http.Request) string {
	cred := &data.Cred{}

	err := data.FromJSON(cred, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(HandlerError{Message: "Invalid JSON"}, w)
		return ""
	}

	if len(cred.URL) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(HandlerError{Message: "Invalid JSON field; 'url'"}, w)
		return ""
	}

	if err = checkURL(cred.URL); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		data.ToJSON(HandlerError{
			Message: fmt.Sprintf("Invalid URL: %s", err.Error()),
		}, w)
		return ""
	}

	return cred.URL
}

// checkURL checks if given URL starts with http or https
func checkURL(URL string) error {
	if !strings.HasPrefix(URL, "http") || !strings.HasPrefix(URL, "https") {
		return fmt.Errorf("URL %s must start with http or https", URL)
	}
	return nil
}
