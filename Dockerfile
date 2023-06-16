
FROM golang:latest

WORKDIR /openai-api-proxy

COPY . /openai-api-proxy

RUN go build -o openaiapiproxy .

FROM scratch

WORKDIR /openai-api-proxy

COPY --from=0 /openaiapiproxy .

CMD ["./openaiapiproxy", "server", "-local-addr", "0.0.0.0:8080", "-db-location", "/data/sqlite3.db"]

EXPOSE 8080
