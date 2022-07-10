package main

import (
	"log"
	"net"

	"github.com/Jat-hom-Wu/ebook/dao"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"
	"github.com/Jat-hom-Wu/ebook/proto"
	"github.com/Jat-hom-Wu/ebook/rpc_service.go"
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

	lis,err := net.Listen("tcp", "127.0.0.1:20022")
	if err != nil{
		log.Fatal("rpc server listen failed,",err)
	}
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &rpc_service.RpcService{})
	if err = s.Serve(lis); err != nil{
		log.Fatal("rpc server serve failed,",err)
	}
}