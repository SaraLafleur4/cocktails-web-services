FROM golang:1.24-alpine3.21 AS builder

WORKDIR /go/src/app
COPY . .
RUN go mod download \
    && go build -ldflags "-s -w" -o /go/bin/main

FROM scratch
COPY --from=builder /go/bin/main /main

EXPOSE 8080

CMD ["/main"]
