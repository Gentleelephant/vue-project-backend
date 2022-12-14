package middleware

import (
	"net/http"

	global2 "github.com/Gentleelephant/vue-project-backend/common/global"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // 先调用c.Next()执行后面的中间件
		// 所有中间件及router处理完毕后从这里开始执行
		// 检查c.Errors中是否有错误
		// 若是自定义的错误则将code、msg返回
		// 若非自定义错误则返回详细错误信息err.Error()
		// 比如save session出错时设置的err
		// 检查一个错误就行
		for _, e := range c.Errors {
			err := e.Err
			if myErr, ok := err.(*global2.CustomError); ok {
				c.JSON(http.StatusOK, global2.Response{
					Code:    myErr.Code,
					Status:  "error",
					Message: myErr.Message,
					Data:    myErr.Data,
				})
				return
			} else {
				c.JSON(500, global2.Response{
					Code:    500,
					Status:  "error",
					Message: err.Error(),
					Data:    nil,
				})
				return
			}
		}
	}
}
