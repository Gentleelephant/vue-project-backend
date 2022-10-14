package group

import (
	"net/http"

	"github.com/Gentleelephant/vue-project-backend/utils"

	"github.com/Gentleelephant/vue-project-backend/config"
	"github.com/Gentleelephant/vue-project-backend/handler"
	"github.com/Gentleelephant/vue-project-backend/model"
	"github.com/Gentleelephant/vue-project-backend/model/global"
	"github.com/gin-gonic/gin"
)

func AccountRegister(c *gin.Context) {
	var err error
	var registerAccount model.RegisterAccount
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()
	err = c.ShouldBindJSON(&registerAccount)
	if err != nil {
		err = global.NewCustomError(global.ErrDataBind, err)
		return
	}
	account, err := handler.AddAccount(config.DB, &model.Account{
		Userid:   utils.RandomUserID(),
		Username: registerAccount.Username,
		Password: registerAccount.Password,
		Gender:   registerAccount.Gender,
		Email:    registerAccount.Email,
	})
	if err != nil {
		err = global.NewCustomError(global.ErrUserRegister, err)
		return
	}
	c.JSON(http.StatusOK, global.Response{
		Code:    200,
		Status:  "success",
		Message: "注册成功",
		Data:    account.IgnorePassword(),
	})
}

func GetRouter(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			_ = c.Error(err)
		}
	}()

	if err != nil {
		err = global.NewCustomError(global.ErrDataBind, err)
	}

}
