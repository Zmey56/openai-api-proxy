FROM golang:latest AS gobuild

WORKDIR /openai-api-proxy

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o openai-api-proxy .

FROM debian:bookworm

RUN apt-get update && apt-get install -y ca-certificates && rm -fR /var/lib/apt/lists/*

COPY --from=gobuild /openai-api-proxy/openai-api-proxy /openai-api-proxy

VOLUME /data/

EXPOSE 8080

ENTRYPOINT ["/openai-api-proxy", "server"]

CMD ["-local-addr", "0.0.0.0:8080", "-db-location", "/data/sqlite3.db"]








