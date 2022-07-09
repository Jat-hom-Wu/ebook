package router

import (
	"net/http"

	"github.com/Jat-hom-Wu/ebook/router/api"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	r.Use(AccessJsMiddleware()) //cors
	group := r.Group("/ebook")
	{
		group.POST("/login", api.LoginHandler)
		group.POST("/register", api.RegisterHandler)
		group.GET("/test", testHandler)
	}

	authGroup := r.Group("/ebook")
	authGroup.Use(api.CheckToken)
	{
		authGroup.GET("/info", api.InfoHandler)
		authGroup.POST("/uploadfile", api.UploadFileHandler)
	}

	listGroup := r.Group("/ebook/list")
	listGroup.Use(api.CheckToken)
	{
		listGroup.GET("/info", api.ListInfoHandler)
		listGroup.POST("/add", api.ListAddHandler)
		listGroup.PUT("/update/:id", api.ListUpdateHandler)
		listGroup.DELETE("/delete/:id", api.ListRemoveHandler)
	}
}

//for webbench
func testHandler(c *gin.Context){
	c.JSON(200,"ok")
	
}

func ServerRun(addr string){
	r.Run(addr)
}

func AccessJsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	   w := c.Writer
	   // 处理js-ajax跨域问题
	   w.Header().Set("Access-Control-Allow-Origin", "http://159.75.2.47:9998") //允许访问所有域
	   w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, POST, GET, PUT, DELETE")
	   w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	   w.Header().Add("Access-Control-Allow-Headers", "Access-Token")
	   w.Header().Set("Access-Control-Allow-Credentials", "true")
	//    w.Header().Add("X-Content-Type-Options", "nosniff")

	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Origin", "http://159.75.2.47:9998")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization") //自定义 Header
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.AbortWithStatus(http.StatusNoContent)
	}
	   c.Next()
	}
 }