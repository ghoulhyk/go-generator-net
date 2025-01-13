package respDecoder

import (
	"encoding/json"
	"github.com/ghoulhyk/go-generator-net/request/resp"
)

type Decoder func(data []byte, result resp.IResp) error

var JsonDecoder Decoder = func(data []byte, result resp.IResp) error {
	if err := json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}
