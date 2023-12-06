FROM golang:1.21.4-bookworm

WORKDIR app

RUN apt install -y libc6

COPY src/main main

CMD [ "main"]