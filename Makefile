all: build

build:
	make -C api

deploy-local: run

run:
	make -C api docker-build
	@docker-compose up -d

stop:
	@docker-compose stop
	@docker-compose rm

test:
	make -C api test

.PHONY: deploy-local
