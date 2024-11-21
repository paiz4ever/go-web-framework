package api

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/paiz4ever/go-web-framework/api/v1"
)

func InitRoutes(r *gin.Engine) {
	v1.InitRoutes(r.Group("/api/v1"))
}
