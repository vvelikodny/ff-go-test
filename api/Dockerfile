# build stage
FROM golang:1.12 AS builder

WORKDIR /go/github.com/vvelikodny/ff-go-test/api
COPY . /go/github.com/vvelikodny/ff-go-test/api

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN make
RUN pwd
RUN ls -l

# run stage
FROM alpine
COPY --from=builder /go/github.com/vvelikodny/ff-go-test/api/bin/api /api

ENTRYPOINT ["/api"]
