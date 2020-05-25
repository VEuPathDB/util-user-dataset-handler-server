FROM alpine:latest AS go

# Prepare environment
ENV GOROOT="/usr/lib/go" \
    GOPATH="/go" \
    PATH="/go/bin:$PATH" \
    GOTAR="go1.14.2.linux-amd64.tar.gz"

# Install Dependencies
RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin \
    && wget https://dl.google.com/go/${GOTAR} \
    && tar -C /usr/lib -xvf ${GOTAR} \
    && rm ${GOTAR} \
    && apk add --no-cache nodejs npm make \
    && npm install -g raml2html raml2html-modern-theme

WORKDIR ${GOPATH}/src/github.com/VEuPathDB/util-user-dataset-handler-server

COPY . .
RUN make
