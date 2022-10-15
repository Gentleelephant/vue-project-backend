package serve

import (
	"github.com/Gentleelephant/vue-project-backend/apis"
	"github.com/Gentleelephant/vue-project-backend/config"
	"github.com/Gentleelephant/vue-project-backend/middleware"
	"github.com/gin-gonic/gin"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Start() {

	// 初始化
	config.Initial()
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.Cors())
	//r.Use(middleware.CheckSession())

	apis.Register(r)

	err := r.Run(":12080")
	if err != nil {
		return
	}

}
