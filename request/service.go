package request

import (
	"github.com/ghoulhyk/go-generator-net/request/req"
)

type Service struct {
}

func (receiver Service) Reqs() []req.Req {
	return []req.Req{}
}

func (receiver Service) BaseUrl() string {
	return ""
}

func (receiver Service) BaseUrlFunc() func() string {
	return nil
}

func (receiver Service) ErrorLog(format string, v ...any) {
}
