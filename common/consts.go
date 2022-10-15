package common

import (
	"fmt"
	"time"
)

const (
	Enable  = iota + 1 // 启用
	Disable            // 禁用
	Delete             // 删除
)

func DeletePrefix() string {
	return fmt.Sprintf("delete-%d-", time.Now().Unix())
}
