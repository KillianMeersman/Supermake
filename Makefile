run: build
	./supermake help

build: test
	go mod tidy
	go generate ./...
	go build -o supermake

test:
	go test -v ./...

install: build
	sudo cp -f supermake /usr/local/bin/supermake

uninstall:
	sudo rm -f /usr/local/bin/supermake

clean:
	rm supermake || true.*

.PHONY: run build clean
