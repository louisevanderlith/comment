FROM scratch

COPY cmd/cmd .

EXPOSE 8084

ENTRYPOINT [ "./cmd" ]