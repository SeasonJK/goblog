package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)
var (
	AppMode string
	HttpPort string
	JwtKey string

	Db string
	Dbname string
	DbPort string
	DbHost string
	DbUser string
	DbPassword string

	AccessKey string
	SecretKey string
	Bucket string
	QiniuServer string
)


func init(){
	file, err := ini.Load("config/config.ini")
	if err != nil{
		fmt.Println("配置文件错误，请检查文件路径：", err)
	}
	LoadServer(file)
	LoadDatabase(file)
	LoadQiNiu(file)
}

func LoadServer(file *ini.File)  {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("43ss426yss")

}

func LoadDatabase(file *ini.File){
	Db = file.Section("database").Key("Db").MustString("mysql")
	Dbname = file.Section("database").Key("Dbname").MustString("goblog")
	DbHost = file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("123123")
}

func LoadQiNiu(file *ini.File){
	AccessKey =file.Section("qiniu").Key("AccessKey").String()
	SecretKey =file.Section("qiniu").Key("SecretKey").String()
	Bucket =file.Section("qiniu").Key("Bucket").String()
	QiniuServer =file.Section("qiniu").Key("QiniuServer").String()
}














