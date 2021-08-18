FROM alpine

COPY ./common /tmp/common

WORKDIR /tmp/

RUN chmod +x common