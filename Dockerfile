FROM golang:1.16.6-alpine3.13 as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o gandidyndns .

FROM alpine:3.13.5
COPY --from=builder /build/gandidyndns /app/

RUN addgroup -g 2000 gandi && adduser -D -u 2000 -G gandi gandi
USER gandi

WORKDIR /app
ENTRYPOINT ["./gandidyndns"]
