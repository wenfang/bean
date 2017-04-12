package db

import (
	"nkwangwenfang.com/kit/transport/http/handler"
)

var (
	errSetDB = handler.NewErrorType(20100, "set db error")
)
