package handler

import (
	"testing"

	"github.com/Gentleelephant/vue-project-backend/common"
	"github.com/Gentleelephant/vue-project-backend/model"
)

func TestCreateRoleSQL(t *testing.T) {

	var role = &model.Role{
		RoleName:   "role3",
		RoleDesc:   "role3",
		RoleStatus: common.Enable,
		Accounts: []*model.Account{
			{
				Aid: 1,
			},
		},
	}

	DB.Model(&model.Role{}).Create(role)

}
