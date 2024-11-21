package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	SUCCESS = 0
	ERROR   = 7
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "ok", c)
}

func OkWithMessage(c *gin.Context, message string) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(SUCCESS, data, "ok", c)
}

func OkWithDetail(c *gin.Context, data interface{}, message string) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "fail", c)
}

func FailWithMessage(c *gin.Context, message string) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetail(c *gin.Context, data interface{}, message string) {
	Result(ERROR, data, message, c)
}

func NoAuth(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		ERROR,
		nil,
		message,
	})
}
