VERSION=$(shell git describe --tags || "untagged")

build: bin/config.yml bin/server bin/static-content/index.html

bin/config.yml:
	@mkdir -p bin
	@cp config.yml bin/config.yml

bin/server:
	@go build -o bin/server --ldflags="-X 'main.version=${VERSION}'" cmd/server.go

bin/static-content/index.html: openapi.yml
	@sed "s/__VERSION__/${VERSION}" > openapi_tmp.yml
	@redoc-cli bundle openapi_tmp.yml
	@mkdir -p bin/static-content
	@mv redoc-static.html bin/static-content/index.html
	@rm openapi_tmp.yml


# Git push prep tasks
git-push:
	@go test ./...

git-pre-commit: docs docs docs docs
	@git add docs/api.html docs/index.html docs/config.html docs/commands.html

docs/api.html: openapi.yml
	@redoc-cli bundle openapi.yml --output docs/api.html
docs/index.html: readme.adoc
	@asciidoctor -b html5 -D docs/ -o index.html -r pygments.rb readme.adoc
docs/config.html: extras/readme/config-file.adoc
	@asciidoctor -b html5 -D docs/ -o config.html -r pygments.rb extras/readme/config-file.adoc
docs/commands.html: extras/readme/commands.adoc
	@asciidoctor -b html5 -D docs/ -o commands.html -r pygments.rb extras/readme/commands.adoc
