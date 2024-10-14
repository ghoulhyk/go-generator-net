package main

import (
	"flag"
	"github.com/ghoulhyk/go-generator-net/gen"
	"github.com/ghoulhyk/go-generator-net/loader"
	"github.com/ghoulhyk/go-generator-net/types"
	"github.com/ghoulhyk/go-generator-net/util/fileutil"
	"github.com/ghoulhyk/go-generator-net/util/jsonutil"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var err error

	if len(os.Args) < 2 {
		panic("请选择命令")
	}

	cmd := os.Args[1]
	switch cmd {
	case "generate":
		{
			err = generate()
			if err != nil {
				panic(err)
			}
		}
	default:
		panic("未知命令")
	}
	return

}

func generate() error {
	println("gen")
	var err error

	if len(os.Args) < 3 {
		panic("请选择源目录")
	}

	srcDir := os.Args[2]
	if !fileutil.IsDir(srcDir) {
		panic("源目录不存在或不是目录")
	}

	srcDir, err = filepath.Abs(srcDir)
	if err != nil {
		return err
	}

	flagSet := flag.NewFlagSet("", flag.ExitOnError)
	targetDirPtr := flagSet.String("target", "", "生成目录")
	err = flagSet.Parse(os.Args[3:])
	if err != nil {
		return err
	}
	if targetDirPtr == nil || strings.TrimSpace(*targetDirPtr) == "" {
		return errors.New("请指定生成目录")
	}
	targetDir := *targetDirPtr
	targetDir, err = filepath.Abs(targetDir)
	if err != nil {
		return err
	}

	if fileutil.Exist(targetDir) {
		if !fileutil.IsDir(targetDir) {
			return errors.New("生成目标不是目录")
		}
	} else {
		err = os.Mkdir(targetDir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	decodeResultJson, err := loader.Load(srcDir)
	if err != nil {
		return err
	}

	//decodeResultJson :=
	//	"[{\"StructName\":\"Common\",\"TmplPkgPath\":\"yb/app/common/internal/remote/templ\",\"GetBaseUrlMode\":1,\"BaseUrl\":\"xxxxxxxx\",\"Reqs\":[{\"Name\":\"SendSmsPOST\",\"Path\":\"sms/p/:pathArg2/send/:pathArg1\",\"Method\":\"POST\",\"ReturnType\":{\"Ident\":\"fddRespModel.BaseResp[fddRespModel.VerifyUrlData]\",\"PkgPathList\":[\"yb/app/common/internal/model/dto/remoteRespModel/fddRespModel\"]},\"HeaderArgs\":null,\"QueryArgs\":[{\"ParaName\":\"mobile\",\"Type\":5,\"ReqName\":\"mobile\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"sign\",\"Type\":5,\"ReqName\":\"sign\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"templId\",\"Type\":8,\"ReqName\":\"templId\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"templArgs\",\"Type\":1,\"ReqName\":\"templArgs\",\"CustomType\":{\"Ident\":\"map[string]interface {}\",\"PkgPathList\":[\"\"]},\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ReqName\":\"channel\",\"Value\":{},\"_getValWay\":\"Static\"}],\"PathArgs\":[{\"ParaName\":\"pathArg1\",\"Type\":5,\"ReqName\":\"pathArg1\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"pathArg2\",\"Type\":5,\"ReqName\":\"pathArg2\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"}],\"BodyArgs\":[{\"ParaName\":\"bodyArg1\",\"Type\":3,\"ReqName\":\"bodyArg1\",\"CustomType\":null,\"PtrType\":true,\"_getValWay\":\"Dynamic\"}],\"PathFormatter\":\"\",\"ReplacePathArgNames\":null},{\"Name\":\"SendSmsGET\",\"Path\":\"sms/g/send\",\"Method\":\"GET\",\"ReturnType\":{\"Ident\":\"fddRespModel.BaseResp[fddRespModel.VerifyUrlData]\",\"PkgPathList\":[\"yb/app/common/internal/model/dto/remoteRespModel/fddRespModel\"]},\"HeaderArgs\":null,\"QueryArgs\":[{\"ParaName\":\"mobile\",\"Type\":5,\"ReqName\":\"mobile\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"sign\",\"Type\":5,\"ReqName\":\"sign\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"templId\",\"Type\":8,\"ReqName\":\"templId\",\"CustomType\":null,\"PtrType\":false,\"_getValWay\":\"Dynamic\"},{\"ParaName\":\"templArgs\",\"Type\":1,\"ReqName\":\"templArgs\",\"CustomType\":{\"Ident\":\"map[string]interface {}\",\"PkgPathList\":[\"\"]},\"PtrType\":false,\"_getValWay\":\"Dynamic\"}],\"PathArgs\":[],\"BodyArgs\":[],\"PathFormatter\":\"\",\"ReplacePathArgNames\":null}]}]\n"
	var decodeResult []types.Service
	err = jsonutil.FromJson(decodeResultJson, &decodeResult)
	if err != nil {
		return err
	}

	err = gen.GenerateServ(targetDir, decodeResult)
	if err != nil {
		return err
	}
	err = gen.GenerateClient(targetDir, decodeResult)
	if err != nil {
		return err
	}
	return nil
}

//go run -mod=mod entgo.io/ent/cmd/ent generate ./entschema --target ./ent --feature sql/execquery --feature sql/modifier --template ./../../../../_/entTemplate/v1/selectorExt.tmpl --template ./../../../../_/entTemplate/v1/queryPaginationExt.tmpl --template ./../../../../_/entTemplate/v1/edgesExt.tmpl --template ./../../../../_/entTemplate/v1/funcsExt.tmpl --template ./../../../../_/entTemplate/v1/sliceFieldsExt.tmpl --template ./../../../../_/entTemplate/v1/sliceEdgesQueryExt.tmpl
