FROM golang:alpine as builder
RUN apk update && apk add git && apk add binutils && apk add ca-certificates
RUN adduser -D -g '' appuser
COPY . $GOPATH/src/github.com/kostiamol/centerms
WORKDIR $GOPATH/src/github.com/kostiamol/centerms
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo \
    -ldflags="-w -s" -o $GOPATH/bin/centerms .
WORKDIR $GOPATH/bin
RUN strip --strip-unneeded centerms

FROM scratch
EXPOSE 8090 8080 8070
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/bin/centerms /go/bin/centerms
USER appuser
ENTRYPOINT ["/go/bin/centerms"]
