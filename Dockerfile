FROM golang AS builder

WORKDIR /app

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s"

FROM scratch

COPY --from=builder /app/fate_no_bot /go/bin/fate_no_bot

ENV TELEGRAM_TOKEN YOUR_TOKEN_FROM_BOTFATHER

ENTRYPOINT ["/go/bin/fate_no_bot"]
