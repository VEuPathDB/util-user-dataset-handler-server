VERSION = $(shell git describe --tags 2>/dev/null || echo "untagged")
COMMIT  = $(shell (git rev-parse HEAD 2>/dev/null || echo "uncommitted") | cut -b 1-12)
DATE    = $(shell date +"%Y-%m-%dT%T%z")
FILES   = $(shell find . -name '*.go')

define go_build
	env CGO_ENABLED=0 GOOS=linux \
	  go build \
	    -o bin/$(1) \
	    --ldflags="\
	      -X 'github.com/VEuPathDB/util-exporter-server/pkg/meta.version=$(VERSION)' \
	      -X 'github.com/VEuPathDB/util-exporter-server/pkg/meta.commit=$(COMMIT)' \
	      -X 'github.com/VEuPathDB/util-exporter-server/pkg/meta.buildDate=$(DATE)' \
	      -s -w" \
	    cmd/$(1)/main.go
endef

# Just build the server binary & api docs
.PHONY: build
build: bin/server bin/static-content/index.html

# Compile all binaries
.PHONY: build-all
build-all: bin/server bin/static-content/index.html bin/gen-config bin/check-config

# Build all release packages
.PHONY: gh-release
travis: bin/server-${GH_TAG}.tar.gz bin/check-config-${GH_TAG}.tar.gz bin/gen-config-${GH_TAG}.tar.gz

# Pre-push code testing
.PHONY: git-push
git-push:
	@go test ./...

# Pre-commit doc generation
.PHONY: git-pre-commit
git-pre-commit: docs/api.html docs/index.html
	@git add docs/api.html docs/index.html

bin/server: $(FILES)
	$(call go_build,server)

bin/check-config: $(FILES)
	$(call go_build,check-config)

bin/gen-config: $(FILES)
	$(call go_build,gen-config)

bin/server-${GH_TAG}.tar.gz: bin/server bin/static-content/index.html
	@cd bin && tar -czf server-${GH_TAG}.tar.gz server static-content && cd ..

bin/gen-config-${GH_TAG}.tar.gz: bin/gen-config
	@cd bin && tar -czf gen-config-${GH_TAG}.tar.gz gen-config && cd ..

bin/check-config-${GH_TAG}.tar.gz: bin/check-config
	@cd bin && tar -czf check-config-${GH_TAG}.tar.gz check-config && cd ..

bin/static-content/index.html: docs/api.html
	@mkdir -p bin/static-content
	@cp docs/api.html bin/static-content/api.html

docs/api.html: api.raml
	@raml2html --theme raml2html-modern-theme api.raml > docs/api.html

docs/index.html: readme.adoc
	@asciidoctor -b html5 -D docs/ -o index.html -r pygments.rb readme.adoc
