FROM golang:1.12-alpine AS base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh gcc musl-dev && \
    go get github.com/go-delve/delve/cmd/dlv && \
    go get github.com/cespare/reflex
WORKDIR /api
ADD . /api
RUN go build -o main


# ---- DEV ----
FROM base AS dev
ENTRYPOINT reflex -r '\.go$' -s -- sh -c 'go build && dlv debug -l 0.0.0.0:2345 --headless --api-version=2 --log=true'

# ---- PROD ----
FROM alpine AS prod
WORKDIR /api
COPY --from=base /api/main /api/
ENTRYPOINT ./main

