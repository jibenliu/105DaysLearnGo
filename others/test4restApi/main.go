package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func main() {
	router := SetupRouter()
	router.Run(":8881")
}

// 在gin中注册路由，并写了两个接口样例
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// GET接口
	router.GET("/hello", Hello)
	// POST接口
	router.POST("/auth/mail", AuthMail)
	return router
}

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data": map[string]interface{}{
			"id": ctx.Query("id"),
		},
	})
}
func AuthMail(ctx *gin.Context) {
	type Params struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var params Params
	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data": map[string]interface{}{
			"accessToken": "Bearer xxx",
			"email":       params.Email,
		},
	})
}

// 单元测试 go test .
// 压力测试 go test -bench BenchmarkHelloWorld
// 压力测试 go test -bench BenchmarkAuthMail
// 压力测试 go test -bench .
// 压力测试 go test -v -run="none" -bench="BenchmarkHelloWorld" -benchtime=10s benchMain_test.go