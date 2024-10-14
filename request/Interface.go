package request

import (
	"ghoulhyk/go-generator-net/request/req"
)

// Interface
// BaseUrl、BaseUrlFunc 二选一
type Interface interface {
	Reqs() []req.Req
	BaseUrl() string
	BaseUrlFunc() func() string
	ErrorLog(format string, v ...any)
}
