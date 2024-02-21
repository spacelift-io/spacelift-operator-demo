FROM golang:1.21.7 as builder
WORKDIR /build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o demo .

FROM scratch
USER 1000
COPY --from=builder /build/demo /bin/demo
ENTRYPOINT ["/bin/demo"]



