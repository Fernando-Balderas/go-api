# go-api
---

### Execution
```bash
go run main.go
```

### Build then execute
```bash
go build main.go
./main
```

### Consume endpoints
```bash
curl -G http://localhost:8080/
# {"message": "home called"}
curl -G http://localhost:8080/articles
# [{"ID":"1","Title":"Hello","desc":"Article Description","content":"Article Content"},{"ID":"2","Title":"Hello 2","desc":"Article Description","content":"Article Content"}]
```