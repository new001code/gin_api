package api

import (
	"gin_api/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DemoApi struct{}

func (demoApi *DemoApi) Hello(c *gin.Context) {
	res := &responses.ApiResponse{}
	c.JSON(http.StatusOK, res.SuccessWithDataAndMessage(nil, "hello world"))
}
