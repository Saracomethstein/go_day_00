all: run

run:
	@go run cmd/go_day_00/main.go 

test:
	@go test ./...