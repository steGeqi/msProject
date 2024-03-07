package model

import (
	"msProject.com/common"
)

//var (
//NoLegalMobile = errs.NewError(2001, "手机号不合法")
//)

const (
	Success             common.BusinessCode = 200
	LoginMobileNotLegal common.BusinessCode = 2001 //手机号不合法
)
