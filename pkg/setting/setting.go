package setting

import (
	"fmt"

	"github.com/go-ini/ini"
)

var (
	MysqlDB string
	MysqlUser string
	MysqlPassword string
	MysqlPort string
	Port string
	Addr string
	LogPath string
	RedisAddr string
	RedisDB int
	RedisDB2 int	//store token
	PicPath string
)

func InitSetting() error{
	cfg,err := ini.Load("/home/lighthouse/ebook/config/app.ini")
	if err != nil{
		return fmt.Errorf("ini file failed,%v",err)
	}
	MysqlUser = cfg.Section("mysql").Key("user").String()
	MysqlPassword = cfg.Section("mysql").Key("password").String()
	MysqlDB = cfg.Section("mysql").Key("database").String()
	MysqlPort = cfg.Section("mysql").Key("port").String()
	Port = cfg.Section("server").Key("http_port").String()
	Addr = cfg.Section("server").Key("addr").String()
	LogPath = cfg.Section("log").Key("path").String()
	RedisAddr = cfg.Section("redis").Key("addr").String()
	RedisDB,_ = cfg.Section("redis").Key("db").Int()
	RedisDB2,_ = cfg.Section("redis").Key("db2").Int()
	PicPath = cfg.Section("app").Key("pic_path").String()
	return nil
}