FROM golang:1.14 as builder
WORKDIR /go/src/fake-stackoverflow
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -v -o fake-stackoverflow# multistage build
FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN mkdir /templates
COPY --from=builder /go/src/fake-stackoverflow/fake-stackoverflow /fake-stackoverflow
EXPOSE 8080
CMD ["/fake-stackoverflow"]