package group

import (
	"github.com/Gentleelephant/vue-project-backend/config"
	"github.com/Gentleelephant/vue-project-backend/handler"
	"github.com/Gentleelephant/vue-project-backend/model"
	"github.com/Gentleelephant/vue-project-backend/model/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type registerAccount struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
}

func AccountRegister(c *gin.Context) {
	var account registerAccount
	err := c.ShouldBindJSON(&account)
	if err != nil {
		c.JSON(http.StatusOK, global.Response{
			Code:    1002,
			Status:  "",
			Message: "注册失败",
			Data:    nil,
		})
		return
	}
	err = handler.AddAccount(config.DB, &model.Account{
		Username: account.Username,
		Password: account.Password,
		Gender:   account.Gender,
		Email:    account.Email,
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, global.Response{
		Code:    200,
		Status:  "success",
		Message: "注册成功",
		Data:    nil,
	})
}
