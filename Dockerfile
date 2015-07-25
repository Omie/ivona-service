FROM scratch
COPY ivona-service /
ENTRYPOINT ["/ivona-service"]
