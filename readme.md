### 示例

#### 请求结构定义
```go
package templ

import (
	"github.com/ghoulhyk/go-generator-net/request"
	"github.com/ghoulhyk/go-generator-net/request/arg"
	"github.com/ghoulhyk/go-generator-net/request/req"
	"github.com/pkg/errors"
	"hykServer/pkg/logutil"
)

type BaseResp[T any] struct {
	Code    uint32 `json:"code"`
	Msg     string `json:"msg"`
	Data    T      `json:"data"`
	Success bool   `json:"success"`
}

func (receiver BaseResp[T]) IsSuccess() bool {
	return receiver.Success
}

func (receiver BaseResp[T]) ErrInfo() error {
	return errors.New(receiver.Msg)
}

type Wxpusher struct {
	request.Service
}

func (receiver Wxpusher) Reqs() []req.Req {
	return []req.Req{
		req.POST("SendMsg", "send/message").
			SetBody([]arg.Arg{
				arg.String("appToken"),
				arg.String("content"),
				arg.String("summary"),
				arg.Uint8("contentType"),
				arg.Custom("topicIds", []uint32{}),
				arg.Custom("uids", []string{}),
				arg.String("url"),
				arg.Bool("verifyPay"),
			}).
			SetReturnType(BaseResp[any]{}),
	}
}

func (receiver Wxpusher) BaseUrl() string {
	return "https://wxpusher.zjiecode.com/api/"
}

func (receiver Wxpusher) ErrorLog(format string, v ...any) {
	logutil.Errorf(format, v...)
}

```
#### 生成代码
`//go:generate go run -mod=mod github.com/ghoulhyk/go-generator-net generate ./templ --target ./remoteServ`
