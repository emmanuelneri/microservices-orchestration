FROM golang:1.14 AS api-build
ADD . /build
WORKDIR /build
RUN go build -o api .

FROM alpine:3.12
WORKDIR /app
COPY --from=api-build /build .
CMD ["/api"]