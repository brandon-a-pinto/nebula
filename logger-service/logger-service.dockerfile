FROM alpine:latest

RUN mkdir /app

COPY ./logger /app

CMD [ "/app/logger" ]
