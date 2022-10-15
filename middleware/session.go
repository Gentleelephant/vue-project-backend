package middleware

import (
	"strings"

	"github.com/Gentleelephant/vue-project-backend/common/global"

	"github.com/Gentleelephant/vue-project-backend/config"
	"github.com/Gentleelephant/vue-project-backend/handler"
	"github.com/Gentleelephant/vue-project-backend/model"
	"github.com/gin-gonic/gin"
)

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {

		var err error
		defer func() {
			if err != nil {
				_ = c.Error(err)
			}
		}()

		path := c.Request.URL.Path
		method := c.Request.Method
		if strings.HasPrefix("/api/login", path) && method == "POST" {
			c.Next()
			return
		}
		if strings.HasPrefix("/api/register", path) && method == "POST" {
			c.Next()
			return
		}

		if strings.HasPrefix("/router", path) && method == "GET" {
			c.Next()
			return
		}

		if strings.HasPrefix("/ping", path) && method == "GET" {
			c.Next()
			return
		}

		sid, err := c.Cookie("sessionId")
		if err != nil {
			err = global.NewCustomError(global.ErrSessionNotExist, err)
			c.Abort()
			return
		}
		var userData model.UserData
		err = handler.GetSession(c.Request.Context(), config.Rdb, sid, &userData)
		if err != nil {
			err = global.NewCustomError(global.ErrSessionTimeout, err)
			c.Abort()
			return
		}

	}
}
