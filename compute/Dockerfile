FROM golang:1-alpine as builder

WORKDIR /src/compute

COPY go.mod ./

RUN go mod download

COPY . .

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN go build -o main .

FROM scratch

COPY --from=builder /src/compute/main /usr/local/bin/main
ENTRYPOINT ["/usr/local/bin/main"]