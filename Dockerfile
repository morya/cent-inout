FROM golang:1.22 as builder

WORKDIR /gocode
ADD . /gocode

RUN go build -o /app

FROM alpine:3.21

COPY --from=builder /app /app

CMD ["/app"]
