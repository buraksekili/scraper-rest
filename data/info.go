package data

import (
	"encoding/json"
	"io"

	"golang.org/x/net/html"
)

type Images []*PageInfo

type PageInfo struct {
	Image string `json:"image"`
}

// Cred is credential for http.Request.Body.
// Includes URL that is going to be fetched.
type Cred struct {
	URL string `json:"url"`
}

// FromJSON decodes the JSON from given io.Reader, e.g., r.Body
func (c *Cred) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c)
}

func GetImages(images Images, node *html.Node) Images {
	if node.Type == html.ElementNode && (node.Data == "img" || node.Data == "script") {
		for _, a := range node.Attr {
			if a.Key == "src" {
				images = append(images, &PageInfo{Image: a.Val})
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		images = GetImages(images, c)
	}

	return images
}
