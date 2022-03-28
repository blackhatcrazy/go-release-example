FROM scratch
COPY go-release-example /usr/bin/go-release-example
ENTRYPOINT ["/usr/bin/go-release-example"]
