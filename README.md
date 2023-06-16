# openai-api-proxy

To play with it, use the following command:

```bash
go run . server -db db.sqlite3 -openai-token sk-YOUROPENAITOKEN
```

Database is not used right now

After that you can invoke requests to OpenAI like

```bash
curl http://localhost:8080/openai/v1/completions \
  -H "Content-Type: application/json" \
  --user login_4:password4 \
  -d '{
    "model": "text-davinci-003",
    "prompt": "Say this is a test",
    "max_tokens": 7,
    "temperature": 0
  }' -v
```



```bash
curl http://localhost:8080/openai/v1/completions \
  -H "Content-Type: application/json" \
  --user login_1:password1 \
  -d '{
    "model": "text-davinci-003",
    "prompt": "Say this is a test",
    "max_tokens": 7,
    "temperature": 0
  }' -v
```

Or 

```bash
curl http://localhost:8080/openai/v1/images/generations \
  -H "Content-Type: application/json" \
  --user login_1:password1 \
  -d '{
    "prompt": "The emblem for the channel in telegrams on p2p cryptocurrency arbitration",
    "n": 2,
    "size": "1024x1024"
  }' -v
```

```bash
curl http://localhost:8080/openai/v1/chat/completions \
  -H "Content-Type: application/json" \
  --user login_1:password1 \
  -d '{
    "model": "gpt-3.5-turbo",                                                                                         
    "messages": [{"role": "user", "content": "Hello!"}]
  }'  -v
```

To see the full requests and responses, use `-log-level=trace` flag.

```bash
go run . -log-level=trace server -db db.sqlite3 -openai-token sk-YOUROPENAITOKEN
```

TTo create new DB with test users

```bash
go run . initdb -add-test-users
```

Docker
```bash
docker  exec openai-api-proxy /openai-api-proxy/openai-api-proxy initdb -db-location /data/sqlite3.db -add-test-users 

docker run --rm -e OPENAI_TOKEN=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 -p 8080:8080 -v openai-api-proxy-volume:/data --name openai-api-proxy openai-api-proxy

docker run --rm -e OPENAI_TOKEN=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 -p 8080:8080 openai-api-proxy

docker build -t openai-api-proxy .  
```
