package rpc_service

import (
	"context"

	pb "github.com/Jat-hom-Wu/ebook/proto"
	"github.com/Jat-hom-Wu/ebook/service"
)

//maybe more elegant, independently sql conn in struct, others...
type RpcService struct{}

func (r *RpcService)Login(ctx context.Context, req *pb.Request) (*pb.Reply,error){
	resp,err := service.Login(req.Name, req.Password)
	return &pb.Reply{Code: int64(resp.Code), Msg: resp.Msg},err
}

func (r *RpcService)Register(ctx context.Context, req *pb.Request) (*pb.Reply,error){
	resp,err := service.Register(req.Name, req.Password)
	return &pb.Reply{Code: int64(resp.Code), Msg: resp.Msg},err
}