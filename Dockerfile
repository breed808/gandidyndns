FROM golang:1.17.5-alpine3.15 as builder
RUN addgroup -g 2000 gandi && adduser -D -u 2000 -G gandi gandi
USER gandi
RUN mkdir /home/gandi/build
COPY . /home/gandi/build/
WORKDIR /home/gandi/build
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o gandidyndns .

FROM alpine:3.15.0
COPY --from=builder /home/gandi/build/gandidyndns /app/
RUN addgroup -g 2000 gandi && adduser -D -u 2000 -G gandi gandi
USER gandi

WORKDIR /app
ENTRYPOINT ["./gandidyndns"]
