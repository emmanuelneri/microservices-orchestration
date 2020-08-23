FROM golang:1.14-alpine as builder
RUN apk update && apk add --no-cache git
WORKDIR build
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/receiver ./cmd

FROM scratch
COPY --from=builder /bin/receiver /bin/receiver
ENTRYPOINT ["/bin/receiver"]