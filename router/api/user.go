package api

import (
	"log"
	"net/http"
	"time"

	"github.com/Jat-hom-Wu/ebook/model"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

//为了避免循环引用，所以与auth.go中重复了，可以把User放到一个单独的包中，同时供user.go和auth.go引用
type User struct{
	Name string `json:"user"`
	PassWord string `json:"password"`
}

//todo:判断用户名密码输入长度
func LoginHandler(c *gin.Context){
	u := User{}
	err := c.ShouldBindJSON(&u)
	if err != nil{
		logger.Fatal("log in bind json failed,",err)
		c.String(404, "server bind json failed")
	}
	name := u.Name
	passWord := u.PassWord
	resp,err := service.Login(name, passWord)
	if err != nil{
		logger.Fatal("log in errror,",err)
	}
	if resp.Code == 0{
		claims := service.UserClaims{
			service.User{name, passWord},
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
				Issuer:    "user",
			},
		}
		token,err := service.GenerateToken(jwt.SigningMethodHS256, claims, []byte(service.SingingKey))
		if err != nil{
			logger.Fatal("token sign failed,",err)
		}
		c.SetCookie("token", token, 600, "/", "159.75.2.47", false, false)
		model.RCToken.RedisUpdateUser(name, token, TokenDuration)	//store token in redis
	}
	c.JSON(200, resp)
}

func RegisterHandler(c *gin.Context){
	u := User{}
	err := c.ShouldBindJSON(&u)
	if err != nil{
		log.Println("register bind json failed,",err)
		c.String(404, "server bind json failed")
	}
	name := u.Name
	passWord := u.PassWord
	resp,err := service.Register(name, passWord)
	if err != nil{
		logger.Fatal("register error,",err)
	}
	c.JSON(200, resp)
}

func InfoHandler(c *gin.Context){
	c.String(http.StatusOK, "info interface")
}
