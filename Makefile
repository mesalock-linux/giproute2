all:
	go get -d ./...
	go build -o ip
clean:
	@rm -rf ip
