VERSION = $(shell git describe --tags 2>/dev/null || echo "untagged")
FILES   = $(shell find . -name '*.go')

build: bin/server bin/static-content/index.html

travis: bin/server bin/static-content/index.html
	@cd bin && tar -czf server-${TRAVIS_TAG}.tar.gz server static-content && cd ..

bin/server: $(FILES)
	@go build -o bin/server --ldflags="-X 'main.version=${VERSION}'" cmd/service.go

bin/static-content/index.html: docs/api.html
	@mkdir -p bin/static-content
	@cp docs/api.html bin/static-content/api.html

git-push:
	@go test ./...

git-pre-commit: docs/api.html docs/index.html
	@git add docs/api.html docs/index.html docs/config.html docs/commands.html

docs/api.html: api.raml
	@raml2html --theme raml2html-modern-theme api.raml > docs/api.html

docs/index.html: readme.adoc
	@asciidoctor -b html5 -D docs/ -o index.html -r pygments.rb readme.adoc
