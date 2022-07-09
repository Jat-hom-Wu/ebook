package service

import (
	"github.com/Jat-hom-Wu/ebook/model"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
)

type Response struct{
	Code int	//0 success,  1 fail, 2 server error
	Data interface{}
	Msg string
}

func Login(name,password string) (Response,error){
	resp := Response{}
	//redis find
	redisPW,err := model.RCUser.RedisFindUser(name)
	if err != nil{
		resp.Code = 2
		resp.Msg = "server error"
		return resp,err
	}else if redisPW != ""{
		if redisPW == password{
			resp.Code = 0
			resp.Data = model.UserTable{Name:name,Password:password}
			resp.Msg = "log in success"
			return resp, nil
		}
	}
	user,err := model.FindUser(name)
	if err != nil{
		resp.Code = 2
		resp.Msg = "server error"
		return resp,err
	}
	if user.Password != password || user.Name == ""{
		resp.Code = 1
		resp.Msg = "password or name wrong"
		return resp,nil
	}
	//redis update
	err = model.RCUser.RedisUpdateUser(user.Name,user.Password, 0)
	if err != nil{
		logger.Info("redis update data failed,",err)
	}
	resp.Code = 0
	resp.Data = user
	resp.Msg = "log in success"
	return resp, nil
}

func Register(name,password string)(Response, error){
	resp := Response{}
	r,err := model.RCUser.RedisFindUser(name)
	if err != nil{
		logger.Info("redis find user failed,",err)
	}
	if r != ""{
		resp.Code = 1
		resp.Data = resp
		resp.Msg = "user name already exists"
		return resp,nil
	}
	user,err := model.FindUser(name)
	if err != nil{
		resp.Code = 2
		resp.Msg = "server error"
		return resp,err
	}
	if user.Name != ""{
		resp.Code = 1
		resp.Data = resp
		resp.Msg = "user name already exists"
		return resp,nil
	}
	err = model.AddUser(name, password)
	if err != nil{
		resp.Code = 2
		resp.Msg = "server error"
		return resp,err
	}
	//redis update
	err = model.RCUser.RedisUpdateUser(name,password, 0)
	if err != nil{
		logger.Info("redis update data failed,",err)
	}
	resp.Code = 0
	resp.Msg = "register success"
	return resp,nil
}