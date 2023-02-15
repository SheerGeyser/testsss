#!make
ifneq (,$(wildcard ./.env))
	include .env
	export
endif

load-env:
	sh scripts/env.sh && env

lint:
	cd src && golangci-lint run

fix:
	cd src && golangci-lint run --fix

build:
	sh scripts/build.sh

migrate:
	cd migrations && tern migrate

migrate-down:
	cd migrations && tern migrate --destination -1

migrate-test:
	cd migrations && tern migrate --config tern_test.conf

test:
	cd src && go test -race -cover -coverprofile=../cover.out -coverpkg=./... ./... && \
	go tool cover -func ../cover.out && \
	go tool cover -html=../cover.out -o ../cover.html

start: build
	./bin/secretary-bot start


