FROM alpine:latest

RUN mkdir /app

COPY ./listener /app

CMD [ "/app/listener" ]
