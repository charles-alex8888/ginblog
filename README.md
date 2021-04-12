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
