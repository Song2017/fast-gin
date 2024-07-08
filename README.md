# Taste Gin
This is a quick start based on Gin.
1. swagger gen example codes
2. run Gin server with packages and middlewares

## Guide
1. modify ./resources/openapi.yaml
2. bin/gen_swagger_server.sh
3. source bin/configuration.sh
4. bin/run_gin_swagger.sh # go run tests/main.go
5. access swagger: http://localhost:9000/swagger/index.html

## Todos
1. upgrade shell scripts

## Object hierarchy
- example: tests/test_main/main.go
### request flow
- api: Swagger groups
- service: HTTP path and function 
- domain: business codes related to service
- po: persitence object related to db
- db
```
api -- service -- domain -- po -- db
```
### functions
- init: load OS env and prepare
- pkg: extension methods
- remote: remote services


## Init Golang Package
```
go mod init github.com/Song2017/startup
go mod tidy
export GOPROXY=https://goproxy.io

dlv debug server/main.go
b path:line
```