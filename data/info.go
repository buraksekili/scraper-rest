package data

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type Images []*PageInfo

type PageInfo struct {
	ImageURL  string `json:"image"`
	ImageName string `json:"image_name"`
}

// Cred is credential for http.Request.Body.
// Includes URL that is going to be fetched.
type Cred struct {
	URL string `json:"url"`
}

// GetImages traverses the parsed HTML, returns each images in HTML.
func GetImages(images Images, node *html.Node, resp *http.Response) Images {
	if node.Type == html.ElementNode && (node.Data == "img" || node.Data == "script") {
		for _, a := range node.Attr {
			if a.Key == "src" {
				fullURL, _ := resp.Request.URL.Parse(a.Val)
				if !strings.HasSuffix(fullURL.String(), "js") {
					images = append(images, &PageInfo{ImageURL: fullURL.String(), ImageName: a.Val})
				}
			}
		}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		images = GetImages(images, c, resp)
	}

	return images
}
