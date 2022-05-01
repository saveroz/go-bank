# go-bank
Application To Understands How Banks Are Works

## do migration for TiDB
migrate -path ./internal/tidb/migrations -database "mysql://root@tcp(localhost:4000)/bank" up