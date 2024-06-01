package router

import (
	"gin_api/api"

	"github.com/gin-gonic/gin"
)

func InitRouter(g *gin.Engine) {
	// 路由组
	demoRouterGroup := g.Group("/demo")
	{
		demoApi := &api.DemoApi{}
		demoRouterGroup.GET("hello", demoApi.Hello)
	}
}
