FROM golang:latest

WORKDIR /openai-api-proxy

COPY . .

RUN go build -o openai-api-proxy .

CMD ["./openai-api-proxy", "server", "-local-addr", "0.0.0.0:8080", "-db-location", "/data/sqlite3.db"]

EXPOSE 8080