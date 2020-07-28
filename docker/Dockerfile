FROM golang:1.14

RUN apt-get update && apt-get -y install --no-install-recommends apt-utils 2>&1

WORKDIR /app

COPY . /app

RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/google/wire/cmd/wire

RUN wire

ENTRYPOINT CompileDaemon --build="go build -o /usr/local/bin/pekabeta ." --command=/usr/local/bin/pekabeta