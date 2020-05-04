FROM alpine:latest

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
    && npm install -g redoc-cli

WORKDIR ${GOPATH}/src/github.com/VEuPathDB/util-exporter-server
VOLUME bin/

COPY . .

CMD make
