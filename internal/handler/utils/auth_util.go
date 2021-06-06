package utils

import (
	"fmt"

	m "github.com/lp/music-chart/internal/model"
	"github.com/lp/music-chart/util"
)

const (
	ErrMsg = "%s is required"
)

func RegisterValidation(regReq m.User) (err error) {
	if util.RemoveAllWhiteSpace(regReq.FullName) == "" {
		return fmt.Errorf(ErrMsg, "full name")
	}

	if util.RemoveAllWhiteSpace(regReq.Password) == "" {
		return fmt.Errorf(ErrMsg, "password")
	}

	if util.RemoveAllWhiteSpace(regReq.UserName) == "" {
		return fmt.Errorf(ErrMsg, "username")
	}
	return
}
