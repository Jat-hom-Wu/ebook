package router

import (
	"testing"

	"github.com/Jat-hom-Wu/ebook/dao"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"
)

func TestInitRouter(t *testing.T){
	err := setting.InitSetting()
	if err != nil{
		t.Error(err)
	}
	err = dao.InitMysql()
	if err != nil{
		t.Error(err)
	}
	InitRouter()
	ServerRun(":9999")
}