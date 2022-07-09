package setting

import(
	"testing"
)

func TestInitSetting(t *testing.T){
	//运行该测试函数需要修改路径
	err := InitSetting()
	if err != nil{
		t.Error("failed,",err)
	}
}