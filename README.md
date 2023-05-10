# openai-api-proxy

To play with it, use the following command:

```bash
go run . server -db db.sqlite3 -openai-token sk-YOUROPENAITOKEN
```

Database is not used right now

After that you can invoke requests to OpenAI like

```bash
curl http://localhost:8080/openai/v1/chat/completions \
  -H "Content-Type: application/json" \
  -H "Authorization: test" \
  -v \
  -d '{
    "model": "gpt-3.5-turbo",
    "messages": [{"role": "user", "content": "Hello!"}]
  }' 
```

Or 

```bash
curl http://localhost:8080/openai/v1/images/generations \
  -H "Content-Type: application/json" \
  -H "Authorization: test" \
  -d '{
    "prompt": "A cute baby sea otter",
    "n": 2,
    "size": "1024x1024"
  }' -v
```

To see the full requests and responses, use `-log-level=trace` flag.

```bash
go run . -log-level=trace server -db db.sqlite3 -openai-token sk-YOUROPENAITOKEN
```
