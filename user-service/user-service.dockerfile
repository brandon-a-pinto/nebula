FROM alpine:latest

RUN mkdir /app

COPY ./user /app

CMD [ "/app/user" ]
