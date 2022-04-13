package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/NitorCreations/azure-functions-go-handler/internal/gofunc"
)

func init() {
	err := gofunc.LoadTemplates()
	gofunc.ExitIf(err)
}

func help() {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("    gofunc <command> [parameters]\n\n")
	fmt.Printf("The commands are\n\n")
	fmt.Printf("    init             create a new Go Function App in the current directory\n")
	fmt.Printf("    generate [dir]   generate func handler code starting from optional [dir], defaults to current directory\n")
	fmt.Printf("    version          print version info and exit\n")
	fmt.Printf("    help             show this help\n\n")
	os.Exit(1)
}

func version() {
	if info, ok := debug.ReadBuildInfo(); ok {
		fmt.Println("gofunc", info.Main.Version, info.GoVersion)
	}
	os.Exit(0)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		help()
	}

	switch args[0] {
	case "init":
		initialize()
	case "generate":
		generate(args[1:])
	case "version":
		version()
	default:
		help()
	}
}

func initialize() {
	fmt.Println("gofunc: init")

	args := []string{}
	wd := gofunc.Getwd(args)

	vars := gofunc.NewProjectVars()
	err := gofunc.CreateProject(wd, vars)
	gofunc.ExitIf(err)

	if _, err := os.Stat("go.mod"); os.IsNotExist(err) {
		var cmd *exec.Cmd

		modName := filepath.Base(gofunc.Getmd())
		cmd = exec.Command("go", "mod", "init", modName)
		gofunc.ExecCmd(cmd)

		cmd = exec.Command("go", "get", "-u", "github.com/NitorCreations/azure-functions-go-handler")
		gofunc.ExecCmd(cmd)
	} else {
		fmt.Println("go.mod already exists. Skipped!")
	}

	generate(args)

	fmt.Printf("\nRun `go build handler.go && func start` to launch Go Function App\n\n")
}

func generate(args []string) {
	fmt.Println("gofunc: generate")

	wd := gofunc.Getwd(args)
	vars := gofunc.NewHandlerVars()

	err := filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && strings.HasPrefix(info.Name(), ".") {
			return filepath.SkipDir
		}
		if info.Name() == "function.json" {
			fun := gofunc.CreateFunctionVars(wd, path)
			if fun != nil {
				vars.Functions = append(vars.Functions, fun)
			}
		}
		return nil
	})

	gofunc.ExitIf(err)

	// generate handler
	path, err := gofunc.CreateHandler(wd, vars)
	gofunc.ExitIf(err)

	// format generated code
	cmd := exec.Command("go", "fmt", path)
	gofunc.ExecCmd(cmd)
}
