# syntax = docker.io/docker/dockerfile:experimental
FROM golang as build

WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/app


FROM scratch as release
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]
