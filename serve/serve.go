package serve

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/Gentleelephant/vue-project-backend/middleware"
	"github.com/Gentleelephant/vue-project-backend/model/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Start() {

	r := gin.Default()
	r.Use(middleware.Cors())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/api/login", func(c *gin.Context) {
		var userLogin UserLogin
		hash := sha256.New()
		hash.Write([]byte("admin"))
		sum := hash.Sum(nil)
		password := hex.EncodeToString(sum)
		err := c.ShouldBindJSON(&userLogin)
		fmt.Println(userLogin.Username, password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, global.Response{
				Code:    500,
				Status:  "error",
				Message: "内部错误",
				Data:    nil,
			})
			return
		}
		if userLogin.Username == "admin" && userLogin.Password == password {
			c.JSON(http.StatusOK, global.Response{
				Code:    200,
				Status:  "success",
				Message: "登录成功",
				Data:    userLogin,
			})
		} else {
			c.JSON(http.StatusUnauthorized, global.Response{
				Code:    401,
				Status:  "error",
				Message: "未授权",
				Data:    nil,
			})
		}
	})

	r.Run(":12080")

}
