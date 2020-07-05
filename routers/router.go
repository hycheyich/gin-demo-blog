package routers

import (
	v1 "github.com/EDDYCJY/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTag)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags", v1.EditTag)
		apiv1.DELETE("/tags", v1.DelTag)
	}
	return r
}

