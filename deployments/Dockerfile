### STAGE 1: Build ###
FROM golang:1.20-alpine3.18@sha256:bddf6dfe104ad797f1e64532f269d7a9c8faa5163d91024632e1b6a554555bf7 AS builder
WORKDIR /build
ADD ./go.mod .
COPY . .
RUN go build -o server ./cmd/app/main.go


### STAGE 2: Run ###
FROM alpine:3.18@sha256:48d9183eb12a05c99bcc0bf44a003607b8e941e1d4f41f9ad12bdcc4b5672f86
COPY --from=builder /build/server server
CMD ["./server"]
