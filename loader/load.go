package loader

import (
	"bytes"
	"fmt"
	"ghoulhyk/go-generator-net/loader/templates"
	"ghoulhyk/go-generator-net/request"
	"ghoulhyk/go-generator-net/util/fileutil"
	"github.com/pkg/errors"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"text/template"
)

func Load(srcDir string) (string, error) {
	var servInterface = reflect.TypeOf(struct{ request.Interface }{}).Field(0).Type
	pkgs, err := packages.Load(&packages.Config{
		BuildFlags: []string{},
		Mode:       packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule,
	}, srcDir, servInterface.PkgPath())
	if err != nil {
		return "", err
	}
	templPkg := pkgs[0]
	reqServPkg := pkgs[1]
	if reqServPkg.PkgPath != servInterface.PkgPath() {
		templPkg, reqServPkg = pkgs[1], pkgs[0]
	}

	var names []string
	iface := reqServPkg.Types.Scope().Lookup(servInterface.Name()).Type().Underlying().(*types.Interface)
	for k, v := range templPkg.TypesInfo.Defs {
		typ, ok := v.(*types.TypeName)
		if !ok || !k.IsExported() || !types.Implements(typ.Type(), iface) {
			continue
		}
		spec, ok := k.Obj.Decl.(*ast.TypeSpec)
		if !ok {
			return "", fmt.Errorf("invalid declaration %T for %s", k.Obj.Decl, k.Name)
		}
		if _, ok := spec.Type.(*ast.StructType); !ok {
			return "", fmt.Errorf("invalid spec type %T for %s", spec.Type, k.Name)
		}
		names = append(names, k.Name)
	}

	genMainPath, err := writeMainSrcFile(templPkg.PkgPath, names, srcDir)
	if err != nil {
		return "", err
	}
	defer os.RemoveAll(genMainPath)

	return gorun(genMainPath, nil)
}

func writeMainSrcFile(tmplPath string, servNames []string, srcDir string) (string, error) {
	var err error
	genMainPath := filepath.Join(srcDir, "..", ".___netreqGen___")
	if fileutil.Exist(genMainPath) {
		err = os.RemoveAll(genMainPath)
		if err != nil {
			return "", err
		}
	}
	err = os.Mkdir(genMainPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	tmpl, err := template.ParseFS(templates.Fs, "*.tmpl")
	if err != nil {
		return "", err
	}

	genMainFilePath := filepath.Join(genMainPath, "main.go")
	genMainFile, err := os.OpenFile(genMainFilePath, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer genMainFile.Close()

	err = tmpl.Execute(genMainFile, map[string]any{
		"TmplPkg":   tmplPath,
		"ServNames": servNames,
	})
	if err != nil {
		return "", err
	}
	return genMainPath, nil
}

func gorun(target string, buildFlags []string) (string, error) {
	s, err := gocmd("run", target, buildFlags)
	if err != nil {
		return "", fmt.Errorf("netReq/load: %s", err)
	}
	return s, nil
}

// goCmd runs a go command and returns its output.
func gocmd(command, target string, buildFlags []string) (string, error) {
	args := []string{command}
	args = append(args, buildFlags...)
	args = append(args, target)
	cmd := exec.Command("go", args...)
	stderr := bytes.NewBuffer(nil)
	stdout := bytes.NewBuffer(nil)
	cmd.Stderr = stderr
	cmd.Stdout = stdout
	if err := cmd.Run(); err != nil {
		return "", errors.New(stderr.String())
	}
	return stdout.String(), nil
}
