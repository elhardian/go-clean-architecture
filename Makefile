.PHONY: run setup

run:
	@go run cmd/main.go

setup:
	@go mod download