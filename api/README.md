# API 

### Build 
- build application `go build -o api-main`
- build image `docker build -t microservices-orchestration/api .`

### Run
-  Run application ``docker run -it -p 8080:8080 microservices-orchestration/api``

### 
``curl -d "{\"number\": \"123\"}" -X POST http://localhost:8080``