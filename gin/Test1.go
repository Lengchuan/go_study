package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	router := gin.Default()

	router.POST("/test", func(context *gin.Context) {
		res := make(map[string]interface{}, 2)
		res["status"] = 200
		res["msg"] = "testtest"
		res["id"] = "1234567"
		res["name"] = "test"
		context.JSON(http.StatusOK, res)
	})

	router.POST("/test2", func(context *gin.Context) {
		var user User
		context.BindJSON(&user)
		res := make(map[string]interface{}, 2)
		res["status"] = 200
		res["msg"] = "test2test2"
		res["id"] = user.Id
		context.JSON(http.StatusOK, res)
	})

	router.POST("/test3", func(context *gin.Context) {
		res := make(map[string]interface{}, 2)
		res["status"] = 200
		res["msg"] = "testtest"
		context.JSON(http.StatusOK, res)
	})

	router.Run(":8080")
}
