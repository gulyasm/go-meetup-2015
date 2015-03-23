default: clean test
	go build -o calculator cmd/calculator/main.go
	go build -o hello cmd/hello/main.go

test: 
	go test ./...

clean: 
	rm -f calculator
	rm -f hello
