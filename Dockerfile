FROM alpine:latest

COPY comment .
COPY conf conf

EXPOSE 8084

ENTRYPOINT [ "./comment" ]