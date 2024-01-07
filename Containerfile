LABEL author="Matthew Stobbs <matthew@stobbs.ca>"

FROM docker.io/library/golang:1.21 AS builder

WORKDIR /usr/src/app
COPY . .
RUN CGO_ENABLED=0 go build -v -o /usr/local/bin/app -ldflags '-s -w' ./...

FROM scratch AS final

COPY --from=builder /usr/local/bin/app /gohttpcheck
ENTRYPOINT ["/gohttpcheck"]
