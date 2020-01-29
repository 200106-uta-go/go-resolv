include .env

hello:
	echo "Running $(PROJECT_NAME)..."

run: hello
	go run main.go

all: run