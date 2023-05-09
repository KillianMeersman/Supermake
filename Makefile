run: build
	./supermake test

build:
	go generate ./...
	go build -o supermake

clean:
	rm supermake || true
	rm internal/parser/supermakegrammar_* || true
	rm internal/parser/SuperMakeGrammar.tokens internal/parser/SuperMakeGrammarLexer.*
	rm -r internal/parser/.antlr

.PHONY:
