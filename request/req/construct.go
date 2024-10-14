package req

import (
	"github.com/ghoulhyk/go-generator-net/types/const"
)

func GET(name string, path string) Req {
	return Req{
		Name:   name,
		Path:   path,
		Method: _const.GET,
	}
}

func POST(name string, path string) Req {
	return Req{
		Name:   name,
		Path:   path,
		Method: _const.POST,
	}
}
