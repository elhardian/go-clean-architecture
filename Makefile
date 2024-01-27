.PHONY: run setup
ENV ?= local
include libs/config/$(ENV).env
export $(shell sed 's/=.*//' libs/config/$(ENV).env)

run:
	@go run cmd/main.go

setup: go-mod-download migrate

go-mod-download:
	@go mod download

migrate:
	@goose -dir migrations $(DATABASE_DRIVER) "$(DATABASE_CONNECTION_STRING)" up