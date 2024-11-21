package gromore

import (
	"github.com/gin-gonic/gin"
	"github.com/paiz4ever/go-web-framework/pkg/response"
	"github.com/paiz4ever/go-web-framework/service/gromore"
)

type SegmentApi struct {
	gromore.GromoreService
}

type SegmentCreateBody struct {
	CommonBody
	Name     string `json:"name" binding:"required"`
	Data     string `json:"data" binding:"required"`
	AdUnitID int    `json:"ad_unit_id" binding:"required"`
}

func (s *SegmentApi) Create(c *gin.Context) {
	var body SegmentCreateBody
	if err := c.ShouldBindJSON(&body); err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	s.GenerateSign(body.CsjUserID)
	response.Ok(c)
}
