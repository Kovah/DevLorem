FROM alpine:latest
COPY devlorem /
ENTRYPOINT ["/devlorem"]
