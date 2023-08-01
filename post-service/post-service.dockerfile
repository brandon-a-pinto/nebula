FROM alpine:latest

RUN mkdir /app

COPY ./post /app

CMD [ "/app/post" ]
