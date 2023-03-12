# syntax=docker/dockerfile:experimental

# Step 1: build golang binary
FROM golang:1.20 as builder
WORKDIR /opt/app
COPY go.* ./
RUN go mod download
COPY . .
# RUN CGO_ENABLED=0 go build -o apiserver
RUN go build -o apiserver

# Step 2: copy binary from step1
FROM ubuntu:latest
# FROM scratch
WORKDIR /opt/app
COPY --from=builder /opt/app/apiserver ./apiserver
CMD ["/opt/app/apiserver"]
