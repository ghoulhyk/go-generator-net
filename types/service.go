package types

import (
	"fmt"
	"github.com/samber/lo"
	"path"
	"strings"
)

type Service struct {
	StructName     string
	TmplPkgPath    string
	GetBaseUrlMode uint8  // 1:BaseUrl; 2:BaseUrlFunc
	BaseUrl        string // 当TemplGetBaseUrlMode=1时，获取的BaseUrl

	Reqs []Req
}

func (receiver Service) Imports() string {
	imports := []string{
		"_import_tmpl " + receiver.TmplPkgPath,
		"_import_fmt fmt",
		"_import_url net/url",
		"_import_http net/http",
		"_import_json encoding/json",
		"_import_errors errors",
		"_import_bytes bytes",
		"_import_ioutil io/ioutil",
	}
	getArgPathListFunc := func(args []Arg) []string {
		var pathList []string
		for _, arg := range args {
			if dynamicArg, ok := arg.(DynamicArg); ok {
				pathList = append(pathList, dynamicArg.PkgPathList()...)
			}
		}
		return pathList
	}
	for _, req := range receiver.Reqs {
		imports = append(imports, req.ReturnType.PkgPathList...)
		imports = append(imports, getArgPathListFunc(req.QueryArgs)...)
		imports = append(imports, getArgPathListFunc(req.PathArgs)...)
		imports = append(imports, getArgPathListFunc(req.BodyArgs)...)
		imports = append(imports, getArgPathListFunc(req.HeaderArgs)...)
	}
	imports = lo.Uniq(imports)
	imports = lo.WithoutEmpty(imports)
	imports = lo.Map(imports, func(item string, _ int) string {
		if strings.Contains(item, " ") {
			item = strings.ReplaceAll(item, " ", " \"")
			item += "\""
		} else {
			item = fmt.Sprintf("\"%s\"", item)
		}
		return item
	})
	return strings.Join(imports, "\n")
}

func (receiver Service) TmplPkgName() string {
	pkgPath := receiver.TmplPkgPath
	return path.Base(pkgPath)
}

func (receiver *Service) Preprocess() error {
	for i := range receiver.Reqs {
		err := receiver.Reqs[i].Preprocess()
		if err != nil {
			return err
		}
	}
	return nil
}
