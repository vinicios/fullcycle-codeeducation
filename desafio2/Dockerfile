FROM golang:alpine as builder

WORKDIR /app
COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/desafio2 .


FROM scratch

COPY --from=builder /app/desafio2 .

ENTRYPOINT ["/desafio2"]