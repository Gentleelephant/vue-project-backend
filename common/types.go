package common

// QueryAccountParams 保证至少有一个参数就可以
type QueryAccountParams struct {
	Aid int `json:"aid" form:"aid"`

	Userid string `json:"userid" form:"userid"`

	Username string `json:"username" form:"username"`
}
