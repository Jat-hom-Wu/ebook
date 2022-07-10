package main

import (
	"fmt"
	// "context"
	"github.com/Jat-hom-Wu/ebook/dao"
	"github.com/Jat-hom-Wu/ebook/global"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"
	"github.com/Jat-hom-Wu/ebook/proto"
	"github.com/Jat-hom-Wu/ebook/router"
	"google.golang.org/grpc"
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
	GrpcClientInit()
	router.InitRouter()
	router.ServerRun(":9999")
}

func GrpcClientInit(){
	conn,err := grpc.Dial("127.0.0.1:20022", grpc.WithInsecure())
	if err != nil{
			logger.Fatal("rpc client dial failed,",err)
			fmt.Println("rpc client dial failed,",err)
	}
	// defer conn.Close()
	client := proto.NewUserServiceClient(conn)
	global.RpcClient = client

	//test
	// reply,err := global.RpcClient.Login(context.Background(), &proto.Request{Name: "xiaoming", Password: "123"})
	// if err != nil{
	// 	logger.Fatal("rpc client call failed,",err)
	// 	fmt.Println("rpc client call failed,",err)
	// }
	// fmt.Println("code:", reply.Code)
	// fmt.Println("msg:", reply.Msg)
}