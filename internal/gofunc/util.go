package gofunc

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/mod/modfile"
)

func GetModName() string {
	modFile, err := os.Open("go.mod")
	ExitIf(err)
	defer modFile.Close()
	modBytes, _ := ioutil.ReadAll(modFile)
	modName := modfile.ModulePath(modBytes)
	return modName
}

func Getmd() string {
	dir, err := os.Getwd()
	ExitIf(err)
	return dir
}

func Getwd(args []string) string {
	dir := Getmd()
	if len(args) > 0 {
		if !filepath.IsAbs(args[0]) {
			dir = filepath.Clean(dir + "/" + args[0])
		} else {
			dir = args[0]
		}
	}
	return dir
}

func ExecCmd(cmd *exec.Cmd) {
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	fmt.Print(out.String())
	if err != nil {
		os.Exit(1)
	}
}

func ExitIf(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
