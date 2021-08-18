FROM alpine

COPY ./common_api /tmp/common_api

WORKDIR /tmp/

RUN chmod +x common_api