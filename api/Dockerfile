# build stage
FROM golang:1.12 AS builder

WORKDIR $GOPATH/github.com/vvelikodny/ff-go-test/api
COPY . $GOPATH/github.com/vvelikodny/ff-go-test/api

ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
RUN make api

# run stage
FROM alpine
COPY --from=builder ./bin/api /api

ENTRYPOINT ["/api"]
