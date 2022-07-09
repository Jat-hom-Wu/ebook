package api

import (
	"fmt"
	"path"

	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/pkg/setting"
	"github.com/gin-gonic/gin"
)

func UploadFileHandler(c *gin.Context) {
	contentType := c.Request.Header.Get("Content-Type")
	f, err := c.FormFile("avatar")

	if err != nil {
		logger.Fatal("server get file failed,", err)
		c.String(404, "server get form file failed")
	}
	fileName, _ := c.Get("name")
	name, _ := fileName.(string) //不需要判断错误，一定是string
	suffix := path.Ext(f.Filename)
	if suffix != ".png" && suffix != ".jpg" && suffix != ".jpeg"{
		c.String(200, "Unsupported picture format")
	}
	f.Filename = name
	fmt.Println("name:",fileName)
	dst := setting.PicPath+"/"+f.Filename+suffix
	fmt.Println("dst:",dst)
	c.SaveUploadedFile(f, dst)
	c.String(200, "success:"+contentType+","+dst)
}
