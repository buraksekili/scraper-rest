package data

type ImageURL string
type Images []*ImageURL

type PageInfo struct {
	URL string
	Images Images
	Title string
}

