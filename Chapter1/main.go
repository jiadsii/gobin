package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func main() {
	//users:=[]User{{ID: 123,Name: "liu"},{ID: 456,Name: "Du"}}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "user",
			"desc": "welcome to my world",
		})
	})
	//http://127.0.0.1:8080/users/10?name=abc
	r.GET("/users/:id", func(context *gin.Context) {
		id := context.Param("id")
		fmt.Println("hello gin , my id" + id)
		name := context.Query("name")
		context.JSON(http.StatusOK, gin.H{
			"name": name,
			"desc": "hello gin my id" + id,
		})
	})
	//http://127.0.0.1:8080/slice?media=bbc&&media=abc
	r.GET("/slice", func(context *gin.Context) {
		context.JSON(http.StatusOK, context.QueryArray("media"))
	})
	//http://127.0.0.1:8080/map?ids[0]=123&ids[1]=456
	r.GET("/map", func(context *gin.Context) {
		context.JSON(http.StatusOK, context.QueryMap("ids"))
	})
	r.POST("/users", func(context *gin.Context) {
		//创建一个用户
		//name:=context.PostForm("name")
		//wechat:=context.PostForm("wechat")
		//context.JSON(http.StatusOK,gin.H{
		//	"name":name,
		//	"wechat":wechat,
		//})
		var user User
		err := context.BindJSON(&user)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	})
	r.GET("/users", func(context *gin.Context) {
		//获取全部用户
	})
	r.PUT("/user:id", func(context *gin.Context) {
		//更新用户id为xx的信息
	})
	r.DELETE("/user/:id", func(context *gin.Context) {
		//删除用户
	})
	r.Run(":8080")
}
