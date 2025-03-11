ARG BUILDER=golang:1.23-alpine3.21
ARG RUNNER=alpine:3.21

FROM ${BUILDER} AS builder

WORKDIR /workspace

COPY . .

RUN apk --no-cache add gcc musl-dev

RUN go mod download \
  && go mod verify

RUN go build -v -o /usr/local/bin/hellotest main.go

FROM ${RUNNER}

WORKDIR /home/apps

RUN apk --no-cache add ca-certificates curl sqlite \
  && update-ca-certificates

COPY --from=builder /usr/local/bin/hellotest .

ENTRYPOINT ["./hellotest"]
