FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY bin/mailerApp .
COPY templates ./templates

CMD ["./mailerApp"]