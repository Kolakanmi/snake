format:
	goimports -l ./

run:
	go run main.go

build:
	go build -o bin/snake main.go
	go build -o snake main.go

compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/snake-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/snake-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/snake-freebsd-386 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/snake-windows main.go

start:
	goimports -l ./
	go build -o snake main.go
	./snake
