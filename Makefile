build: 
	@go build -o bin/rname

run: build
	@./bin/rname

test: 
	@go test ./... -v
