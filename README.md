[![Build Status](https://travis-ci.com/vvelikodny/ff-go-test.svg?branch=master)](https://travis-ci.com/vvelikodny/ff-go-test)

## Roles

Developer __Vitaly Velikodny__
  * [@vvelikodny](https://github.com/vvelikodny)
  * [vvelikodny@gmail.com](mailto:vvelikodny@gmail.com)  

## Requirements:
  * `go`
  * `docker`
  * `docker-compose`
  * `go get -u golang.org/x/lint/golint`
  
## Deploy process (locally)

Build services and run Docker containers

```bash
make run-env
```

## Demo

Add news

```bash
curl -X POST http://localhost:8080/isgood
  -H "Content-Type: application/json" \
  -d '[{"checkType": "DEVICE","activityType": "SIGNUP","checkSessionKey": "string","activityData": [{"kvpKey": "ip.address","kvpValue": "1.23.45.123","kvpType": "general.string"}]}]'
```
