run: build
	./supermake help

build:
	go generate ./...
	go build -o supermake

install: build
	sudo cp -f supermake /usr/local/bin/supermake

uninstall:
	sudo rm -f /usr/local/bin/supermake

clean:
	rm supermake || true.*

.PHONY: run build clean
