all: build

build:
	make -C api

run-env: build
	make -C api
	@docker-compose up -d

test:
	make -C api test

stop-env:
	@docker-compose stop
	@docker-compose rm

deploy-local: run-env

.PHONY: deploy-local
