# go-request-parser

This code is modified from [Novalagung Golang Basic C4-C6](https://dasarpemrogramangolang.novalagung.com/C-http-error-handling.html). It features request parsing from xml, json, HTML form, query parameter and validate the result to fulfill your business need (e.g user authentication)

### How to run

1. Execute `go build` then `go run main.go`.
2. Send a request (use cURL, Postman, etc) or go to "host:port/form" e.g `localhost:9000/form`
3. If port 9000 (default) is not usable, custom the port with --port or -p e.g `go run main.go -p 8765`

### Request-response example
Given `curl -X GET localhost:9000/user?name=nito&email=nito@mail.com`, the server returns JSON below
```
{
    "name": "nito",
    "email": "nito@mail.com",
    "role": 0
}
```
with status code 200 OK

### Update log

1. Add form with CSRF feature. To use CSRF, make `.env` file, use `.env.example` as template