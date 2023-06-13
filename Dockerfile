FROM golang:latest

WORKDIR /openai-api-proxy

COPY . .

RUN go build -o main .

CMD ["./main"]

EXPOSE 8080