package dao

import (
	"fmt"
	"testing"
	"time"

	"github.com/Jat-hom-Wu/ebook/model"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"
)

func TestAddUser(t *testing.T){
	err := setting.InitSetting()
	if err != nil{
		t.Errorf("%v",err)
	}
	err = InitMysql()
	if err != nil{
		t.Errorf("%v",err)
	}
	// err = model.AddUser("xiaoli","123")
	// if err != nil{
	// 	t.Errorf("%v",err)
	// }
	err = model.UpdateUser("xiaoli12","123456789")
	if err != nil{
		t.Errorf("%v",err)
	}
	user,err := model.FindUser("xiaoli1")
	if err != nil{
		t.Errorf("%v",err)
	}else{
		fmt.Println("user:",user.Name)
	}
}

func TestRedis(t *testing.T){
	err := setting.InitSetting()
	if err != nil{
		t.Errorf("%v",err)
	}
	InitRedis()
	err = model.RCUser.RedisUpdateUser("xiaolong","123",10 * time.Second)
	if err != nil{
		t.Error(err)
	}
	res,err := model.RCUser.RedisFindUser("xiaolong")
	if err != nil{
		t.Error(err)
	}
	fmt.Printf("res:%v\n",res)
}

func TestListTable(t *testing.T){
	err := setting.InitSetting()
	if err != nil{
		t.Errorf("%v",err)
	}
	err = InitMysql()
	if err != nil{
		t.Errorf("%v",err)
	}
	// err = model.ListAdd("xiaoming", "asdfasdf", true)
	// if err != nil{
	// 	t.Error(err)
	// }
	err = model.ListUpdate(false,14)
	if err != nil{
		t.Error(err)
	}
	// err = model.ListDelete(2)
	// if err != nil{
	// 	t.Error(err)
	// }
	val,err := model.ListCheck("xiaoming")
	if err != nil{
		t.Error(err)
	}
	fmt.Println("val:",val)
}