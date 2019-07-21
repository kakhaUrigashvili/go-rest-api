# Rest API Spot Hero - Golang

## Prerequisites

Docker

## Start up

To launch the application run the following command

```
docker-compose up
```

Navigate to http://localhost:8000


## Test

To run the tests run the following command

```
docker-compose run --entrypoint='go test -v ./...' api 
```

## Swagger Docs

Swagger Docs can be found on the main page

http://localhost:8000

## Dockerfile

Dockerfile is in the root folder of the project. The Dockerfile uses multi stage build to have debugging utils
available only in development and ensure small image size in production.


## Metrics

Application metrics are exposes in format consumable by Prometheus.

http://localhost:8000/metrics
