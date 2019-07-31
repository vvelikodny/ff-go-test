include .env

all: $(BINARY_NAME)

clean:
	[ ! -f ./bin/$(BINARY_NAME) ] || rm ./bin/$(BINARY_NAME) || echo "-"

$(BINARY_NAME): clean lint
	echo Building for $(GOOS)/$(GOARCH)
	go build -a -o ./bin/$(BINARY_NAME)

test:
	go test ./...

lint:
	@go get -u golang.org/x/lint/golint
	@go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

docker-build: clean
	@docker build -t $(REGISTRY_IMAGE) .

deploy-local: docker-build

.PHONY: deploy-local
