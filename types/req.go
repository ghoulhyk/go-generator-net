package types

import (
	"fmt"
	"github.com/ghoulhyk/go-generator-net/types/const"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"net/url"
	"regexp"
	"slices"
	"strings"
)

type Req struct {
	Name       string
	Path       string
	Method     _const.ReqMethod
	ReturnType RType
	HeaderArgs Args
	QueryArgs  Args
	PathArgs   Args
	BodyArgs   Args

	PathFormatter       string
	ReplacePathArgNames []string
}

func (receiver Req) args() []Arg {
	return slices.Concat(receiver.HeaderArgs, receiver.PathArgs, receiver.QueryArgs, receiver.BodyArgs)
}

func (receiver Req) ParamArgs() []Arg {
	return lo.Filter(receiver.args(), func(item Arg, _ int) bool {
		_, ok := item.(StaticArg)
		return !ok
	})
}

func (receiver *Req) Preprocess() error {
	receiver.PathFormatter = receiver.Path
	if strings.Contains(receiver.PathFormatter, ":") {
		pathParamMapping := lo.SliceToMap(receiver.PathArgs, func(item Arg) (string, Arg) {
			return item.GetReqName(), item
		})
		pathParamRegex := regexp.MustCompile("(:[a-zA-Z0-9_]+)")
		errAny, _ := lo.TryWithErrorValue(func() error {
			receiver.PathFormatter = pathParamRegex.ReplaceAllStringFunc(receiver.PathFormatter, func(s string) string {
				s = s[1:]
				pathArg, found := pathParamMapping[s]
				if !found {
					panic(errors.Errorf("未找到path参数[%s", s))
				}
				if staticArg, ok := pathArg.(StaticArg); ok {
					return url.PathEscape(staticArg.ValueStr(false))
				} else if dynamicArg, ok := pathArg.(DynamicArg); ok {
					receiver.ReplacePathArgNames = append(receiver.ReplacePathArgNames, dynamicArg.ParaName)
					return dynamicArg.ValueFormatter()
				}
				panic(errors.Errorf("path参数[%s]类型暂不支持", s))
			})
			return nil
		})
		if errAny != nil {
			if err, ok := errAny.(error); ok {
				return err
			} else {
				return errors.Errorf("预处理path参数失败: %v", errAny)
			}
		}
	}
	if len(receiver.QueryArgs) > 0 {
		if !strings.Contains(receiver.PathFormatter, "?") {
			receiver.PathFormatter += "?"
		} else if !strings.HasSuffix(receiver.PathFormatter, "?") {
			receiver.PathFormatter += "&"
		}
		queryItems := lo.Map(receiver.QueryArgs, func(item Arg, _ int) string {
			if staticArg, ok := item.(StaticArg); ok {
				return fmt.Sprintf("%s=%s", staticArg.GetReqName(), url.QueryEscape(staticArg.ValueStr(false)))
			} else if dynamicArg, ok := item.(DynamicArg); ok {
				receiver.ReplacePathArgNames = append(receiver.ReplacePathArgNames, dynamicArg.ParaName)
				return fmt.Sprintf("%s=%s", dynamicArg.GetReqName(), dynamicArg.ValueFormatter())
			}
			return ""
		})
		receiver.PathFormatter += strings.Join(queryItems, "&")
	}
	return nil
}
