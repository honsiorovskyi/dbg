# dbg
Simple docker container that listens for HTTP requests on $LISTEN and resonds with the request dump

## Usage

```
docker build -t dbg .
docker run -d -e LISTEN=:9000 -p 9000:9000 dbg

curl -XPOST http://localhost:9000/ -H 'Some-Header: some-value' -d 'param=data'
```

Response example:
```
POST / HTTP/1.1
Host: localhost:9000
Accept: */*
Content-Length: 10
Content-Type: application/x-www-form-urlencoded
Some-Header: some-value
User-Agent: curl/7.65.1

param=data
```
