# ginblog

# 
> go mod init ginblog
# 安装gin框架
~~~ bash
go env -w GO111MODULE="auto"
go get -u github.com/gin-gonic/gin
go mod tidy
~~~ 

# ini配置文件包
> go get gopkg.in/ini.v1
# gorm包
> go get -u github.com/jinzhu/gorm
# mysql驱动
> _ "github.com/go-sql-driver/mysql"

# 目录框架
- config
- model
- api/v1
- middleware
- routes
- utils
- upload
- web

# 接口测试工具
- ApiPost
## JWT
> go get -u github.com/dgrijalva/jwt-go

## 日志
> go get -u github.com/sirupsen/logrus
## 日志切割
> go get github.com/lestrrat-go/file-rotatelogs
go get github.com/rifflock/lfshook


# 前端
## 
vue create mgt


## 安装yarn
curl -o- -L https://yarnpkg.com/install.sh | bash
yarn add ant-design-vue --dev
yarn add axios
yarn add babel-plugin-import --dev

## 启动
yarn serve
