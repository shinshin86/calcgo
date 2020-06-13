dev:
	go run main.go sum 1 23 && go run main.go sub 20 1 && go run main.go mult 3 2 && go run main.go div 10 5

test:
	go test -v ./cmd

fmt:
	go fmt cmd/
