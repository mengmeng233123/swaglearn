package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)
// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://wang.me

// @contact.name Wang
// @contact.url http://wang.me
// @contact.email 2635773708@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main()  {

	r:=gin.Default()

	//url:=ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	//r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.Run()
	//userGroup=r.Group("users")
	r.Run(":3000")
}