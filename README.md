## Simple Url Shortener

A url shortener that uses a simple hash function to generate a short url from a long url.

it uses the SHA256 hash function to generate a hash from the long url and then uses the first 8 characters of the hash to generate a short url.

it then uses a base58 encoding to convert the hash to a short url.

## usage

```
go run main.go
```

### routes

- /create-short-url
  - method: POST
  - body:  {
    "long_url": "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b36"
}
  - response: { 
    "message": "short url created successfully",
    "short_url": "http://localhost:8000/9Zatkhpi"
 }

- /{short_url}
  - method: GET
  - response
    - redirect to the long url

## packages used
- github.com/gin-gonic/gin
- github.com/go-redis/redis

