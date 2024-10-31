FROM golang:alpine as build
ARG VERSION=1.0.24
ARG UserAgentBase=Pure_FA_OpenMetrics_exporter

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=1 go build -mod=readonly -a -tags 'netgo osusergo static_build' -ldflags="-X 'main.version=v$VERSION' -X 'purestorage/fa-openmetrics-exporter/internal/rest-client.FARestUserAgentBase=$UserAgentBase' -X 'purestorage/fa-openmetrics-exporter/internal/rest-client.UserAgentVersion=$VERSION'" -v -o /usr/local/bin/pure-fa-om-exporter cmd/fa-om-exporter/main.go


# alpine is used here as it seems to be the minimal image that passes quay.io vulnerability scan
FROM alpine
# update ssl packages for CVEs
RUN apk update && apk add --upgrade libcrypto3 libssl3 && rm -rf /var/cache/apk/*
COPY --from=build  /usr/local/bin/pure-fa-om-exporter /pure-fa-om-exporter

# create an empty tokens file for use with volumes if required. You can use a mounted volume to /etc/pure-fa-om-exporter/ to pass the `tokens.yaml` file. File must be named `tokens.yaml`.
RUN mkdir /etc/pure-fa-om-exporter && touch /etc/pure-fa-om-exporter/tokens.yaml

EXPOSE 9490
ENTRYPOINT ["/pure-fa-om-exporter"]
CMD ["--address", "0.0.0.0", "--port", "9490", "--tokens", "/etc/pure-fa-om-exporter/tokens.yaml"]
