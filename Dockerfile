FROM golang:latest AS gobuild

WORKDIR /openai-api-proxy

COPY . .

RUN go build -o openai-api-proxy .

FROM scratch

WORKDIR /openai-api-proxy

COPY --from=gobuild /openai-api-proxy/openai-api-proxy .

CMD ["./openai-api-proxy", "server", "-local-addr", "0.0.0.0:8080", "-db-location", "/data/sqlite3.db"]

EXPOSE 8080








