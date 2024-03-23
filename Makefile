run: build
	@go run .

build: 
	@go build .

test: 
	@go test ./...
