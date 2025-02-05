# Backend Portal
這個專案Golang的後端開發實作

## 指令集

### 初始化專案環境
```bash
go mod init Backend
go mod tidy
go mod vendor
```

### Running the Project
```bash
go run main.go
OR
go run -mod=mod main.go
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

## Directory Structure
```bash
#專案:
├── cmd                      # 編譯
├── conf.d                   # 設定檔 IT 定義的命名規則，P.S. 如果是用 config server 的話可能會改動設定方式，參考就好
│   └── app.yaml             # IT 設定，服務相關 IP 帳密
├── docs                     # 文件
├── service                  # 系統邏輯
│   │   app                  # 對外接口 app 目錄底下的資料夾可自行定義名稱， web, admin, job 為範例
│   │   ├── web              # 對終端用戶
│   │   ├── admin            # 對管理員
│   │   └── job              # 排程設定
│   ├── binder               # di 設定區
│   ├── config               # 設定檔取得
│   ├── constant             # 系統常數、 Enum 定義
│   ├── controller           # core 與 app 的轉換層
│   │   ├── handler          # req resp 基本方法
│   │   └── middleware       # middleware 邏輯區
│   ├── core                 # 業務邏輯 core 目錄底下的資料夾可自行定義名稱， usecase, module 為範例
│   │   ├── usecase          # 使用場景
│   │   └── module           # 底層邏輯
│   ├── errs                 # 錯誤訊息定義
│   ├── model                # 資料結構
│   │   ├── bo               # core 層邏輯使用的結構
│   │   ├── dto              # controller 層的資料使用結構
│   │   └── po               # repository 轉換對外的資料結構
│   ├── repository           # 與資料庫溝通層，如 MySQL, Redis
│   ├── listener             # 接收資料區，如 MQ
│   ├── job                  # 排程
│   ├── manager              # 與三方服務溝通
│   ├── thirdparty           # 三方套件防腐層
│   ├── util                 # 共用小工具
│   ├── version.go           # 紀錄編譯版本與 commit 
│   └── run.go               # 入口點 di, app
└── README.md                # 服務簡介，有目錄連結至 docs
```