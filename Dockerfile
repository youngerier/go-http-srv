FROM golang:alpine3.13 AS builder
WORKDIR /build
RUN go version
COPY . .
RUN ls &&\
    go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/app .
CMD ["./app"]  