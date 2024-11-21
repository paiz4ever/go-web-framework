package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/paiz4ever/go-web-framework/api/v1/gromore"
)

func InitRoutes(r *gin.RouterGroup) {
	gromoreGroup := r.Group("/gromore")
	{
		var segmentApi = new(gromore.SegmentApi)
		gromoreGroup.POST("/create-segment", segmentApi.Create)
	}
}
