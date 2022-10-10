package serve

import (
	"github.com/Gentleelephant/vue-project-backend/middleware"
	"github.com/gin-gonic/gin"
)

func Start() {

	r := gin.Default()
	r.Use(middleware.Cors())
	r.POST("/api/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":12080")

}
