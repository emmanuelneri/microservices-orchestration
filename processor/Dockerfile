FROM golang:1.14 as app-builder
WORKDIR build
ADD . .
# TODO: improve image - RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api
RUN go build -o processor

#TODO change to alpine
FROM golang:1.14
WORKDIR /app
COPY --from=app-builder /go/build .
ENTRYPOINT /app/processor