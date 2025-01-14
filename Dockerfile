FROM golang:1.17 as build-env

WORKDIR /go/src/app
COPY * .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 go build -o /go/bin/app

FROM gcr.io/distroless/static
ENV GIN_MODE=release
COPY ./index.html /index.html
COPY --from=build-env /go/bin/app /
CMD ["/app"]