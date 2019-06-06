FROM golang:alpine AS build
RUN apk add --update --no-cache ca-certificates git
ADD . /src
RUN cd /src && CGO_ENABLED=0 go build -v -o sonos_exporter

FROM scratch
COPY --from=build /src/sonos_exporter /
EXPOSE 1915
ENTRYPOINT ["/sonos_exporter"]
