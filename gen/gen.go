package gen

import (
	"bytes"
	"github.com/ghoulhyk/go-generator-net/gen/templates"
	"github.com/ghoulhyk/go-generator-net/types"
	"github.com/ghoulhyk/go-generator-net/util/fileutil"
	"go/format"
	"golang.org/x/tools/go/packages"
	"os"
	"path/filepath"
	"text/template"
)

const servDirName = "internal"

func GenerateServ(dstPath string, services []types.Service) error {
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"dict":   dict,
			"pascal": pascal,
		}).
		ParseFS(templates.Fs, "serv.tmpl", "req.tmpl")
	if err != nil {
		return err
	}
	servDirPath := filepath.Join(dstPath, servDirName)
	if !fileutil.Exist(servDirPath) {
		err = os.Mkdir(servDirPath, os.ModePerm)
		if err != nil {
			return err
		}
	}
	for _, service := range services {
		err = service.Preprocess()
		if err != nil {
			return err
		}
		srcBuffer := bytes.NewBuffer(nil)
		err = tmpl.ExecuteTemplate(srcBuffer, "service", service)
		if err != nil {
			return err
		}
		buf, err := format.Source(srcBuffer.Bytes())
		if err != nil {
			return err
		}
		servFilePath := filepath.Join(servDirPath, service.StructName+".go")
		err = os.WriteFile(servFilePath, buf, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func GenerateClient(dstPath string, services []types.Service) error {
	pkgs, err := packages.Load(&packages.Config{
		BuildFlags: []string{},
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
	}, filepath.Join(dstPath, servDirName))
	if err != nil {
		return err
	}
	servPkgPath := pkgs[0].PkgPath
	pkgName := filepath.Base(dstPath)
	tmpl, err := template.New("").
		Funcs(template.FuncMap{
			"dict":   dict,
			"pascal": pascal,
		}).
		ParseFS(templates.Fs, "client.tmpl")
	if err != nil {
		return err
	}

	srcBuffer := bytes.NewBuffer(nil)
	err = tmpl.ExecuteTemplate(srcBuffer, "client", map[string]any{
		"PkgName":     pkgName,
		"ServPkgPath": servPkgPath,
		"Services":    services,
	})
	if err != nil {
		return err
	}
	buf, err := format.Source(srcBuffer.Bytes())
	if err != nil {
		return err
	}
	servFilePath := filepath.Join(dstPath, "client.go")
	err = os.WriteFile(servFilePath, buf, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
