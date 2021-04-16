package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	JwtKey     string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误！", err)
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("Server").Key("Appmode").MustString("debug")
	HttpPort = file.Section("Server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("Server").Key("JwtKey").MustString("4cgrhujmu4a89xjv")
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("ginblog123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")
}
