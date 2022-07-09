package api

//后续把jwt改成md5或其他,这里不该使用jwt

import (
	"time"

	"github.com/Jat-hom-Wu/ebook/model"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/service"
	"github.com/gin-gonic/gin"
)

const TokenDuration time.Duration = 10 * time.Minute

func CheckToken(c *gin.Context){
	//get user name
	strToken,err := c.Cookie("token")
	if err != nil{
		logger.Fatal("server get token failed,",err)
		c.Abort()
		c.JSON(200, service.Response{
			Code: 2,
			Msg: "no token",
		})
		return
	}
	token,err := service.ParseToken(strToken, &service.UserClaims{}, []byte(service.SingingKey))
	if err != nil{
		logger.Fatal("parse token failed,",err)
		c.Abort()
		c.JSON(200, service.Response{
			Code: 2,
			Msg: "parse token failed",
		})
		return
	}
	user,ok := token.Claims.(*service.UserClaims) 
	if !ok || !token.Valid{
		c.Abort()
		c.JSON(200, service.Response{
			Code: 2,
			Msg: "token not match",
		})
		return
	}
	name := user.Name

	//judging if user name is in redis
	s,err := model.RCToken.RedisFindUser(name)
	if err != nil{
		logger.Fatal("redis find failed,",err)
		c.Abort()
		c.JSON(200, service.Response{
			Code: 2,
			Msg: "server can't find token",
		})
		return
	}
	if s == ""{
		c.Abort()
		c.JSON(200, service.Response{
			Code: 2,
			Msg: "token not exist",
		})
		return
	}
	c.Set("name", name)
}