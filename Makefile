build:
	go build -o bin/server main.go

build-linux:
	env GOOS=linux GOARCH=amd64 go build -o bin/server main.go