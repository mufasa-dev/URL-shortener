.PHONY: deafult run build test docs clean
# Variables
APP_NAME=url_shortener

# Tasks
default: run_with_docs

run:
	@go run main.go
run_with_docs:
	@swag init
	@go run main.go
build:
	@go build -o ${APP_NAME} main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f ${APP_NAME}
	@rm -rf ./docs