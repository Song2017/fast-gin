# server-spike

## Init Golang Package
```
go mod init server
export GOPROXY=https://goproxy.io
go mod tidy

dlv debug server/main.go
b path:line
```
## Guide
1. modify demo.yaml
2. bin/gen_swagger_server.sh
3. source bin/configuration.sh
4. bin/run_gin_swagger.sh
5. access swagger: http://localhost:9000/swagger/index.html
