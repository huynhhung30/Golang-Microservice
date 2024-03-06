FROM golang:1.22-alpine
EXPOSE 5001

WORKDIR /app

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
# ENTRYPOINT ["/out/main"]  