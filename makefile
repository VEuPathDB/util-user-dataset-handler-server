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
