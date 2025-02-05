:: 第一層
mkdir cmd,conf.d,docs

:: 新增 app.yaml 檔案
echo. > conf.d\app.yaml

:: 多層
mkdir service\internal\app\web
mkdir service\internal\app\admin
mkdir service\internal\app\job
mkdir service\binder
mkdir service\config
mkdir service\constant
mkdir service\controller\handler
mkdir service\controller\middleware
mkdir service\core\usecase
mkdir service\core\module
mkdir service\errs
mkdir service\model\bo
mkdir service\model\dto
mkdir service\model\po
mkdir service\repository
mkdir service\listener
mkdir service\job
mkdir service\manager
mkdir service\thirdparty
mkdir service\util

:: 新增 README.md、gitlab-ci.yml、go.mod、go.sum、run.go、version.go
echo. > README.md
echo. > service\run.go
echo. > service\version.go
