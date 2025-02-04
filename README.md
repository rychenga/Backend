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
├── cmd                   # 編譯
├── conf.d                # 設定檔 IT 定義的命名規則
│   ├── app.yaml          # IT 設定，服務相關 IP 帳密
│   └── config.yaml       # 服務的設定檔
├── docs                  # 文件
├── service               # 系統主要功能
│   ├── internal          # 系統邏輯，不對外
│   │   ├── app           # 對外接口 app 目錄底下的資料夾可自行定義名稱，web, admin, job 為範例
│   │   │   ├── web       # 對終端用戶
│   │   │   ├── admin     # 對管理員
│   │   │   └── job       # 排程設定
│   ├── binder            # di 設定區
│   ├── config            # 設定檔取得
│   ├── constant          # 系統常數、Enum 定義
│   ├── controller        # core 與 app 的轉換層
│   └── handler           # req resp 基本方法

#核心的專案結構：
├── middleware            # middleware 運轉區
├── core                  # 業務邏輯 core 目錄底下的資料夾可自行定義名稱，usecase,module 為範例
│   ├── usecase           # 使用場景
│   ├── module            # 底層邏輯
│   ├── errs              # 錯誤訊息定義
│   ├── model             # 資料結構
│   │   ├── bo            # core 層邏輯使用的結構
│   │   ├── dto           # repository 層的資料傳遞結構
│   │   └── repository    # controller 對應對外的資料結構
│   ├── repository        # 與資料庫通訊，如 PostgreSQL, MySQL, Redis
│   ├── listener          # 接收數據，如 MQ
│   ├── manager           # 負責服務流量
│   ├── thirdparty        # 第三方套件的應用
│   └── util              # 共用小工具
└── vendor                # gomod vendor 模式下相依套件擺放區，不是 vendor mode 的不可用有
```