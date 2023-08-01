FROM alpine:latest

RUN mkdir /app

COPY ./broker /app

CMD [ "/app/broker" ]
