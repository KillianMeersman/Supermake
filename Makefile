run: build
	./supermake help

build:
	go generate ./...
	go build -o supermake

clean:
	rm supermake || true.*

.PHONY: run build clean
