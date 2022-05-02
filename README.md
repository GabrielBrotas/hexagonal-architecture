
```
    go mod init github.com/GabrielBrotas/hexagonal-architeture

    go test ./
```

```
    apt-get update
    apt-get install sqlite3

    touch sqlite.db
    sqlite3 sqlite.db
    
    create table products(id string, name string, price float, status string);
    .tables
```

Cobra CLI
```bash
go install github.com/spf13/cobra-cli@latest

cobra-cli init
go mod tidy
go run main.go

## add commands
cobra-cli add cli


go run main.go cli
go run main.go cli --help
go run main.go cli -a="create" -n="Product Cli" -p=20.5

## run http server
go run main.go http
```