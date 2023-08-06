FROM golang:1.19 AS build
WORKDIR /go/src
COPY . .

ENV CGO_ENABLED=0
ENV GOARCH=amd64
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o books-service-origin ./cmd/books-service-origin

FROM alpine:3.15 AS runtime
ENV BOOKS_SERVICE_ORIGIN_API_ENDPOINT=:80
ENV BOOKS_SERVICE_ORIGIN_MAINTENANCE_ENDPOINT=:8000
COPY --from=build /go/src/books-service-origin ./
EXPOSE 80/tcp
EXPOSE 8000/tcp
ENTRYPOINT ["./books-service-origin"]
