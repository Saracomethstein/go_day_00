all: run

build:
	@go build -o go_day_00 cmd/go_day_00/main.go

run:
	@go run cmd/go_day_00/main.go 

test:
	@go test ./...

clean:
	@rm ./go_day_00