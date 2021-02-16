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

FROM busybox

WORKDIR /compute/

ARG UID=1000
ARG GID=1000

#RUN addgroup -S -g $GID compute
#RUN adduser -S -u $UID -s /bin/sh -G compute compute

COPY --from=builder /src/compute/main .

#USER compute

RUN touch /tmp/mydata.txt

ENTRYPOINT ["/compute/main"]