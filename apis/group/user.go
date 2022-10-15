package group

import (
	"net/http"

	"github.com/Gentleelephant/vue-project-backend/common"

	global2 "github.com/Gentleelephant/vue-project-backend/common/global"
	"github.com/Gentleelephant/vue-project-backend/common/utils"

	"github.com/Gentleelephant/vue-project-backend/config"
	"github.com/Gentleelephant/vue-project-backend/handler"
	"github.com/Gentleelephant/vue-project-backend/model"
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
		err = global2.NewCustomError(global2.ErrDataBind, err)
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
		err = global2.NewCustomError(global2.ErrUserRegister, err)
		return
	}
	c.JSON(http.StatusOK, global2.Response{
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

	var params common.QueryAccountParams
	err = c.ShouldBindQuery(&params)
	if err != nil {
		err = global2.NewCustomError(global2.ErrDataBind, err)
	}

	menu, err := handler.GetMenuByAccount(config.DB, &params)
	router := model.ArrayToTree(menu)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, global2.Response{
		Code:    200,
		Status:  "success",
		Message: "成功",
		Data:    router,
	})

}
