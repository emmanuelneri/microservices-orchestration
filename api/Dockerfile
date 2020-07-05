FROM golang:1.14 as app-builder
WORKDIR build
ADD . .
# TODO: improve image - RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api
RUN go build -o api

#TODO change to alpine
FROM golang:1.14
ENV KAFKA_BOOTSTRAP_SERVERS "10.1.6.112:9092,10.1.6.114:9092,10.1.6.116:9092"
WORKDIR /app
COPY --from=app-builder /go/build .
ENTRYPOINT /app/api