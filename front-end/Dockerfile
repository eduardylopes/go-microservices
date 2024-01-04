FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY bin/frontApp .
COPY cmd/web/templates cmd/web/templates

CMD ["./frontApp"]