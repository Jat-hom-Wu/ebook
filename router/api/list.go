package api

import (
	"fmt"
	"strconv"

	"github.com/Jat-hom-Wu/ebook/model"
	"github.com/Jat-hom-Wu/ebook/pkg/logger"
	"github.com/Jat-hom-Wu/ebook/service"
	"github.com/gin-gonic/gin"
)


func ListAddHandler(c *gin.Context){
	obj := model.ListTable{}
	c.Bind(&obj)
	if obj.Title == ""{
		return
	}
	name,_ := c.Get("name")
	strName,_ := name.(string)
	err := service.AddList(obj.Title, strName)
	if err != nil{
		logger.Fatal("add list failed,",err)
	}
	c.String(200,"")
}

func ListUpdateHandler(c *gin.Context){
	strId,ok := c.Params.Get("id")
	if !ok{
		c.String(404, "url id error")
	}
	list := model.ListTable{}
	err := c.BindJSON(&list)
	if err != nil{
		c.String(404, "server bind json failed")
		logger.Fatal("server bind json failed",err)
		return
	}
	fmt.Println("status:", list.Status)
	id,_ := strconv.Atoi(strId)
	err = service.UpdateList(list.Status, id)
	if err != nil{
		logger.Fatal("update list failed,",err)
		c.String(404,"update list failed")
	}
	c.String(200, "update list success")
}

func ListRemoveHandler(c *gin.Context){
	id,ok := c.Params.Get("id")
	if !ok{
		logger.Fatal("listdelete get id failed,")
		c.String(404,"get id failed")
	}
	err := service.DeleteList(id)
	if err != nil{
		logger.Fatal("list delete failed",err)
		c.String(200,"delete failed")
	}
	c.String(200,"delete ok")
}

func ListInfoHandler(c *gin.Context){
	name,_ := c.Get("name")
	strName,_ := name.(string)
	resp,err := service.GetList(strName)
	if err != nil{
		logger.Fatal("get list info failed,",err)
		c.JSON(200, []int{})
	}
	c.JSON(200, resp)

}