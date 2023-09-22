run:
	go run .

build-all: build-linux build-mac build-win

build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/msis.linux .

build-mac:
	GOOS=darwin GOARCH=amd64 go build -o ./build/msis.mac .

build-win:
	GOOS=windows GOARCH=amd64 go build -o ./build/ .

clean:
	rm -rf ./build/*
