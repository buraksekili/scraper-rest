package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/buraksekili/scraper-rest/data"
	"golang.org/x/net/html"
)

// Info holds HTTP handler for getting page information.
type Info struct {
	l *log.Logger
}

// GetNewInfo creates handler with logger for HTTP requests
func GetNewInfo(l *log.Logger) *Info {
	return &Info{l: l}
}

// ServeHTTP is a must for implementing http.Handler for Info
// https://golang.org/pkg/net/http/#Handler
func (i *Info) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		url := getURL(w, r)
		fmt.Println("URL", url)
		if url == "" || len(url) == 0 {
			http.Error(w, "Invalid URL provided", http.StatusBadRequest)
			return
		}

		images, err := fetchURL(url)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		encoder := json.NewEncoder(w)
		err = encoder.Encode(images)
		if err != nil {
			http.Error(w, "Unable to marshal json", http.StatusInternalServerError)
		}

		return
	}
	http.Error(w, fmt.Sprintf("Expect method POST, got %v", r.Method), http.StatusMethodNotAllowed)
}

// fetchURL fetches the images of given URL
func fetchURL(URL string) (data.Images, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s: %s\n", URL, resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("parsing html error: %s: %s\n", URL, resp.Status)
	}

	return data.GetImages(nil, doc), nil
}

// getURL extracts URL from JSON taken from r.Body
func getURL(w http.ResponseWriter, r *http.Request) string {
	cred := &data.Cred{}

	err := cred.FromJSON(r.Body)
	if err != nil {
		http.Error(w, "Couldn't decode given JSON", http.StatusBadRequest)
		return ""
	}

	if len(cred.URL) < 1 {
		http.Error(w, "Invalid JSON is given", http.StatusBadRequest)
		return ""
	}

	return cred.URL
}
