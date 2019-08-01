all: build

build:
	make -C api

run: build
	make -C api docker-build
	@docker-compose up -d

test:
	make -C api test

stop:
	@docker-compose stop
	@docker-compose rm

deploy-local: run

.PHONY: deploy-local
