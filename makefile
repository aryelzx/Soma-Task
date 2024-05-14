.PHONY: default run build 

APP_NAME=soma-task

default: run

run:
	@go run cmd/soma-task/soma-task.go
build:
	@go build -o $(APP_NAME) cmd/soma-task/soma-task.go