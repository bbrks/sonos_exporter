FROM golang:alpine AS build
RUN apk add --update --no-cache git
ADD . /src
RUN cd /src && CGO_ENABLED=0 go build -o sonos_exporter

FROM scratch
COPY --from=build /src/sonos_exporter /
EXPOSE 1915
ENTRYPOINT ["/sonos_exporter"]
