run: build
	bin/supermake help

build: test
	go mod tidy
	go mod vendor
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
