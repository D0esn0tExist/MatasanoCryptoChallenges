APP=matasano

## setup: setup go modules.
setup:
	@go mod tidy

## build: build the app.
test:
	# @echo "Testing..."
	@go test ./... -v

## help: prints this help message
help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: test clean help