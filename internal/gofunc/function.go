package gofunc

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/NitorCreations/azure-functions-go-handler/pkg/function"
)

func CreateFunctionVars(wd string, path string) *FunctionVars {
	md := Getmd()
	modName := GetModName()

	config, err := function.NewConfigFromFile(path)
	ExitIf(err)

	if config.Excluded {
		return nil
	}

	var funName string = "Handle"
	if config.EntryPoint != "" {
		funName = config.EntryPoint
	}

	cfgPath := strings.TrimPrefix(path, wd+"/")
	pkgName := filepath.Dir(strings.TrimPrefix(path, md+"/"))
	pkgAlias := filepath.Base(pkgName)

	return &FunctionVars{
		Name:       pkgAlias,
		Reference:  fmt.Sprintf("%s.%s", pkgAlias, funName),
		ConfigPath: cfgPath,
		ImportPath: fmt.Sprintf("%s/%s", modName, pkgName),
	}
}
