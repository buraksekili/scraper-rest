# scraper-rest
![carbon](https://user-images.githubusercontent.com/32663655/106610344-bd842600-6577-11eb-8193-3132d8e96154.png)

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
The server accepts POST requests from port :3000

```shell script
$ curl localhost:3000 -d '{"url":"https://golang.org/pkg/net/http/"}'
```
It returns JSON response as follows;
```JSON
[{"image":"https://golang.org/lib/godoc/images/go-logo-blue.svg","image_name":"/lib/godoc/images/go-logo-blue.svg"},{"image":"https://golang.org/lib/godoc/images/footer-gopher.jpg","image_name":"/lib/godoc/images/footer-gopher.jpg"}]
```

If `jq` is installed, you can pipe the JSON result to get better view.
```shell script
$ curl localhost:3000 -d '{"url":"https://golang.org/pkg/net/http/"}' | jq

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
![carbon (1)](https://user-images.githubusercontent.com/32663655/106610401-cecd3280-6577-11eb-95d9-f9b689b41f94.png)


## Acknowledgments
Image scraper is based on Chapter 5 of the [The Go Programming Language](https://www.gopl.io/)

### License
[MIT](https://github.com/buraksekili/scraper-rest/blob/master/LICENSE)
