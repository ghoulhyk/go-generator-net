package req

import (
	"github.com/ghoulhyk/go-generator-net/request/arg"
	"github.com/ghoulhyk/go-generator-net/request/req/contentType"
	"github.com/ghoulhyk/go-generator-net/request/resp"
	"github.com/ghoulhyk/go-generator-net/request/resp/respDecoder"
	"github.com/ghoulhyk/go-generator-net/types/const"
)

type Req struct {
	Name       string
	Path       string
	Method     _const.ReqMethod
	ReturnType resp.IResp
	HeaderArgs []arg.Arg
	QueryArgs  []arg.Arg
	PathArgs   []arg.Arg
	BodyArgs   []arg.Arg

	// todo
	ContentType contentType.Type
	// todo
	ResponseDecoder respDecoder.Decoder
}

func (receiver Req) SetReturnType(returnType resp.IResp) Req {
	receiver.ReturnType = returnType
	return receiver
}

func (receiver Req) SetContentType(typeVal contentType.Type) Req {
	receiver.ContentType = typeVal
	return receiver
}

func (receiver Req) SetResultDecoder(decoder respDecoder.Decoder) Req {
	receiver.ResponseDecoder = decoder
	return receiver
}

// region params

func (receiver Req) AddHeader(header arg.Arg) Req {
	receiver.HeaderArgs = append(receiver.HeaderArgs, header)
	return receiver
}

func (receiver Req) SetHeaders(headers []arg.Arg) Req {
	receiver.HeaderArgs = headers
	return receiver
}

// endregion

// region params

func (receiver Req) AddQuery(query arg.Arg) Req {
	receiver.QueryArgs = append(receiver.QueryArgs, query)
	return receiver
}

func (receiver Req) SetQueries(queries []arg.Arg) Req {
	receiver.QueryArgs = queries
	return receiver
}

// endregion

// region paths

func (receiver Req) AddPath(path arg.Arg) Req {
	receiver.PathArgs = append(receiver.PathArgs, path)
	return receiver
}

func (receiver Req) SetPaths(paths []arg.Arg) Req {
	receiver.PathArgs = paths
	return receiver
}

// endregion

// region body

func (receiver Req) AddBodyItem(bodyItem arg.Arg) Req {
	receiver.BodyArgs = append(receiver.BodyArgs, bodyItem)
	return receiver
}

func (receiver Req) SetBody(body []arg.Arg) Req {
	receiver.BodyArgs = body
	return receiver
}

// endregion
