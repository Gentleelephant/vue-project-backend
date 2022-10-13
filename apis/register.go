package apis

import (
	"github.com/Gentleelephant/vue-project-backend/apis/group"
	"github.com/gin-gonic/gin"
)

func registerUserInterface(g *gin.RouterGroup) {
	g.POST("/login", group.Login)
	g.POST("/register", group.AccountRegister)
}

func registerInterface(g *gin.RouterGroup) {
	g.GET("/ping", group.Ping)
}

func Register(engine *gin.Engine) {
	gapi := engine.Group("/api")
	g := engine.Group("")
	registerInterface(g)
	registerUserInterface(gapi)
}
