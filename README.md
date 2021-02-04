# scraper-rest
![scraper-rest](https://user-images.githubusercontent.com/32663655/106951708-2a96e780-6741-11eb-85b0-d4ae47e12d8b.png)

scraper-rest serves simple image scrapper as REST API.

## Installation
Clone this repository to your local environment. 
Then, start the server.

```shell script
$ git clone https://github.com/buraksekili/scraper-rest.git
$ cd scraper-rest
$ go run main.go
```

## Usage
The server accepts POST requests from port :5000

```shell script
$ curl localhost:5000/images -d '{"url":"https://golang.org/pkg/net/http/"}'
```
It returns JSON response as follows;
```JSON
[{"image":"https://golang.org/lib/godoc/images/go-logo-blue.svg","image_name":"/lib/godoc/images/go-logo-blue.svg"},{"image":"https://golang.org/lib/godoc/images/footer-gopher.jpg","image_name":"/lib/godoc/images/footer-gopher.jpg"}]
```

If `jq` is installed, you can pipe the JSON result to get better view.
```shell script
$ curl localhost:5000/images -d '{"url":"https://golang.org/pkg/net/http/"}' | jq

[
  {
    "image": "https://golang.org/lib/godoc/images/go-logo-blue.svg",
    "image_name": "/lib/godoc/images/go-logo-blue.svg"
  },
  {
    "image": "https://golang.org/lib/godoc/images/footer-gopher.jpg",
    "image_name": "/lib/godoc/images/footer-gopher.jpg"
  }
]
```


## Acknowledgments
Image scraper is based on Chapter 5 of the [The Go Programming Language](https://www.gopl.io/)

Highly inspired by [this repo](https://github.com/nicholasjackson/building-microservices-youtube/tree/episode_7/product-api).

### License
[MIT](https://github.com/buraksekili/scraper-rest/blob/master/LICENSE)
