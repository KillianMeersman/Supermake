run: build
	bin/supermake help

vendor:
	go mod tidy
	go mod vendor

build: test
	go generate ./...
	go build -mod vendor -o bin/supermake

test:
	go test -v ./...

install: build
	sudo cp -f bin/supermake /usr/local/bin/supermake
	sudo cp -lf /usr/local/bin/supermake /usr/local/bin/smake

uninstall:
	sudo rm -f /usr/local/bin/supermake

clean:
	rm -rf bin || true.*

.PHONY: run build clean
