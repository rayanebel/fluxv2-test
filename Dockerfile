ARG RUNNER=alpine:3.19
FROM ${RUNNER}

WORKDIR /home/apps
COPY hellotest .

RUN apk --no-cache add ca-certificates curl jq sqlite \
  && update-ca-certificates

ENTRYPOINT ["./hellotest"]
