FROM golang:1.11-alpine as goalpine
WORKDIR /go/src/github.com/eidosam/tortoise
ADD . .
RUN CGO_ENABLED=0 go build -o /tortoise ./cmd/server.go

FROM alpine
RUN apk --no-cache add ca-certificates
EXPOSE 8080
COPY --from=goalpine /tortoise /bin/tortoise
ENTRYPOINT /bin/tortoise
