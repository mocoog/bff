FROM golang:1.17.3-buster

WORKDIR /go/src/work

COPY / /go/src/work/

RUN go mod tidy && \
    go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949 && \
    go install -v golang.org/x/tools/gopls@v0.7.3 && \
    go install -v github.com/go-swagger/go-swagger/cmd/swagger@v0.28.0 && \
    go build ./...

EXPOSE 8080