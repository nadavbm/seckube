# Build the manager binary
FROM golang:1.20-bullseye as builder

# Copy the go source
COPY . /

WORKDIR /

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o seckube main.go

FROM alpine:latest

WORKDIR /

COPY --from=builder /seckube /seckube

CMD /seckube
