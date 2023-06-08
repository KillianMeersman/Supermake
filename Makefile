run: build
	./supermake test

build:
	go generate ./...
	go build -o supermake

clean:
	rm supermake || true.*

.PHONY: run build clean
