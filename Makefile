default: clean test
	go build -o calculator cmd/calculator/main.go
	go build -o hello cmd/hello/main.go

test: 
	go test ./...

cover:
	go test ./add -coverprofile=coverage.out
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out
	rm coverage.out

clean: 
	rm -f calculator
	rm -f hello

doc:
	 godoc -notes "TODO|BUG" github.com/gulyasm/go-meetup-2015/add 
