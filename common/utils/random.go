package utils

import "github.com/duke-git/lancet/v2/random"

const (
	IdLengthWithoutPrefix = 10
	SessionIDLength       = 32
)

func RandomUserID() string {
	return "usr-" + random.RandString(IdLengthWithoutPrefix)
}

func RandomSession() string {
	return random.RandString(SessionIDLength)
}
