FROM golang:1.12-alpine as builder
LABEL maintainer="Ivan Savcic <isavcic@gmail.com>"
WORKDIR $GOPATH/src/github.com/isavcic/komeon
COPY . .
RUN apk update && apk add upx
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-w -s" -o /komeon .
RUN upx --best --overlay=strip /komeon
FROM scratch
COPY --from=builder /komeon /
ENTRYPOINT ["./komeon"]
