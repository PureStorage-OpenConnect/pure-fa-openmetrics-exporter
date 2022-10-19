FROM golang:1.19.2-alpine

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/pure-fa-om-exporter cmd/fa-om-exporter/main.go

EXPOSE 9490
ENTRYPOINT ["/usr/local/bin/pure-fa-om-exporter"]
CMD ["--host", "0.0.0.0", "--port", "9490"]
