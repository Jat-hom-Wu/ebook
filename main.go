package main

import (
	"github.com/Jat-hom-Wu/ebook/dao"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"
	"github.com/Jat-hom-Wu/ebook/router"
)

func main(){
	logger.SetUpLog("develop","/home/lighthouse/ebook/log")
	err := setting.InitSetting()
	if err != nil{
		logger.Fatal(err)
	}
	err = dao.InitMysql()
	dao.InitRedis()
	if err != nil{
		logger.Fatal(err)
	}
	router.InitRouter()
	router.ServerRun(":9999")
}