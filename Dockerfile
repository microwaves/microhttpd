FROM golang:1.11 as builder

WORKDIR /go/src/github.com/microwaves/microhttpd
ADD . .
RUN CGO_ENABLED=0 go build -o /microhttpd -a -tags netgo -ldflags '-w' microhttpd.go

FROM alpine
COPY --from=builder /go/src/github.com/microwaves/microhttpd/www /www
COPY --from=builder /microhttpd /microhttpd

CMD /microhttpd
