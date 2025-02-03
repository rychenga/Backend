# Backend Portal
這個專案Golang的後端開發實作

## 指令集

### 初始化專案環境
```bash
go mod init Backend
go mod tidy
```

### Running the Project
```bash
go run main.go
```

### Building the Project
```bash
go build -o myprogram.exe main.go
或
go build -ldflags "-s -w" -o myprogram.exe
```

## Environment
### Web Framework
- Gin

### Database
- MySql
- PostgreSQL