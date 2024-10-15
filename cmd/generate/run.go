package generate

import (
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

func run(srcDir, dstDir string) error {
	println("gen")
	var err error

	if !fileutil.IsDir(srcDir) {
		panic("源目录不存在或不是目录")
	}

	srcDir, err = filepath.Abs(srcDir)
	if err != nil {
		return err
	}

	if strings.TrimSpace(dstDir) == "" {
		return errors.New("请指定生成目录")
	}
	dstDir, err = filepath.Abs(dstDir)
	if err != nil {
		return err
	}

	if fileutil.Exist(dstDir) {
		if !fileutil.IsDir(dstDir) {
			return errors.New("生成目标不是目录")
		}
	} else {
		err = os.Mkdir(dstDir, os.ModePerm)
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

	err = gen.GenerateServ(dstDir, decodeResult)
	if err != nil {
		return err
	}
	err = gen.GenerateClient(dstDir, decodeResult)
	if err != nil {
		return err
	}
	return nil
}
