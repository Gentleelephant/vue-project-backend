package group

import (
	_ "errors"
	_ "net/http"

	_ "github.com/Gentleelephant/vue-project-backend/common/global"
	_ "github.com/Gentleelephant/vue-project-backend/common/utils"

	"github.com/Gentleelephant/vue-project-backend/config"
	_ "github.com/Gentleelephant/vue-project-backend/handler"

	//"github.com/Gentleelephant/vue-project-backend/handler/service"
	_ "github.com/Gentleelephant/vue-project-backend/model"
	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Ping(c *gin.Context) {
	config.Plog.Info("This is a test")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//func Login(c *gin.Context) {
//	var err error
//	var userLogin UserLogin
//	defer func() {
//		if err != nil {
//			_ = c.Error(err)
//		}
//	}()
//	err = c.ShouldBindJSON(&userLogin)
//	if err != nil {
//		err = global2.NewCustomError(global2.ErrDataBind, err)
//		return
//	}
//	ok, err := service.CheckPassword(config.DB, userLogin.Username, userLogin.Password)
//	if errors.Is(err, gorm.ErrRecordNotFound) {
//		err = global2.NewCustomError(global2.ErrUserNotExist, err)
//		return
//	}
//	if ok {
//		sessionId := utils.RandomSession()
//		account, err := handler.FindAccountByUsername(config.DB, userLogin.Username)
//		account = account.IgnorePassword()
//		if err != nil {
//			return
//		}
//		err = handler.SetSession(c.Request.Context(), config.Rdb, sessionId, account)
//		if err != nil {
//			return
//		}
//		c.JSON(http.StatusOK, global2.Response{
//			Code:    200,
//			Status:  "success",
//			Message: "登录成功",
//			Data: model.UserData{
//				SessionId: sessionId,
//				Account:   *account,
//			},
//		})
//		return
//	}
//	err = global2.ErrPasswordWrong
//}
