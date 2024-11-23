# https://maex.me/2018/02/dont-fear-the-makefile/
run: 
	go build -o ./bin/todo ./cmd/todo
	./bin/todo

test:
	go test ./...
