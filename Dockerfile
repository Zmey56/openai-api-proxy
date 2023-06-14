FROM golang:latest

WORKDIR /openai-api-proxy

COPY . .

RUN go build -o openai-api-proxy .

CMD ["./openai-api-proxy", "server", "-openai-token", "$OPENAI_TOKEN"]

EXPOSE 8080