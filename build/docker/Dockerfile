FROM golang:alpine as build

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=1 go build -a -tags 'netgo osusergo static_build' -v -o /usr/local/bin/pure-fa-om-exporter cmd/fa-om-exporter/main.go

FROM scratch
#FROM busybox:musl
COPY --from=build  /usr/local/bin/pure-fa-om-exporter /pure-fa-om-exporter
EXPOSE 9490
ENTRYPOINT ["/pure-fa-om-exporter"]
CMD ["--address", "0.0.0.0", "--port", "9490"]
