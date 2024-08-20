build:
	@go build -o bin/forminit cmd/app/main.go

test: 
	@go test -v ./...

run: build
	@./bin/forminit