FROM golang:alpine

ADD . /go/src/app
WORKDIR /go/src/app
ENV PORT=8080
RUN go get .
RUN go install

# COPY go.mod .
# COPY go.sum .
# RUN go mod download
# COPY [ "main.go", "models", "httputil", "controllers",  "./"]

# RUN go build -o /app

EXPOSE $PORT

ENTRYPOINT ["/go/bin/app"]
