# go-request-parser

This code is modified from [Novalagung Golang Basic C4-C6](https://dasarpemrogramangolang.novalagung.com/C-http-error-handling.html). It features request parsing from xml, json, HTML form, query parameter and validate the result to fulfill your business need (e.g user authentication)

### How to run

1. Execute `go build` then `go run main.go`.
2. Send a request (use cURL, Postman, etc)

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