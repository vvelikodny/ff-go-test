[![Build Status](https://travis-ci.com/vvelikodny/ff-go-test.svg?branch=master)](https://travis-ci.com/vvelikodny/ff-go-test)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/vvelikodny/ff-go-test)](https://cloud.docker.com/repository/docker/vvelikodny/ff-go-test)

## Roles

Developer __Vitaly Velikodny__
  * [@vvelikodny](https://github.com/vvelikodny)

## Requirements:
  * `docker`
  * `docker-compose`
  
## Deploy process (locally)

Build services and run Docker containers

```bash
make deploy-local
```

## Demo

Add news

```bash
curl -X POST http://localhost:8080/isgood \
  -H "Content-Type: application/json" \
  -d '[{"checkType": "DEVICE","activityType": "SIGNUP","checkSessionKey": "string","activityData": [{"kvpKey": "ip.address","kvpValue": "1.23.45.123","kvpType": "general.string"}]}]'
```
