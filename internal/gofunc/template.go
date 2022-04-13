package gofunc

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

const (
	templatesDir = "templates"
)

var (
	//go:embed templates/*
	files     embed.FS
	templates map[string]*template.Template
)

var (
	handlerTemplate  = "handler.tmpl"
	projectTemplates = map[string]string{
		"HttpTrigger/function.json": "http.function.json.tmpl",
		"HttpTrigger/main.go":       "http.main.go.tmpl",
		".funcignore":               "funcignore.tmpl",
		".gitignore":                "gitignore.tmpl",
		"host.json":                 "host.json.tmpl",
		"local.settings.json":       "local.settings.json.tmpl",
	}
)

func LoadTemplates() error {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	tmplFiles, err := fs.ReadDir(files, templatesDir)
	if err != nil {
		return err
	}

	for _, file := range tmplFiles {
		if file.IsDir() {
			continue
		}

		tmpl, err := template.ParseFS(files, templatesDir+"/"+file.Name())
		if err != nil {
			return err
		}

		templates[tmpl.Name()] = tmpl
	}
	return nil
}

type HandlerVars struct {
	Functions []*FunctionVars
}

type FunctionVars struct {
	Name       string
	Reference  string
	ConfigPath string
	ImportPath string
}

type ProjectVars struct {
	ExecExt string
}

func NewHandlerVars() *HandlerVars {
	return &HandlerVars{
		Functions: []*FunctionVars{},
	}
}

func NewProjectVars() *ProjectVars {
	execExt := ""
	if runtime.GOOS == "windows" {
		execExt = ".exe"
	}

	return &ProjectVars{
		ExecExt: execExt,
	}
}

func CreateHandler(dir string, vars *HandlerVars) (string, error) {
	path := filepath.Join(dir, "handler.go")
	flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC

	file, err := os.OpenFile(path, flags, 0664)
	if err != nil {
		return "", err
	}

	defer file.Close()

	if tmpl, ok := templates[handlerTemplate]; ok {
		return path, tmpl.Execute(file, vars)
	}
	return "", fs.ErrNotExist
}

func CreateProject(dir string, vars *ProjectVars) error {
	dir, err := filepath.Abs(dir)
	if err != nil {
		return err
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return err
		}
	}

	for fp, tp := range projectTemplates {
		path := filepath.Join(dir, fp)
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			fmt.Printf("%s already exists. Skipped!\n", fp)
			continue
		}

		subdir := filepath.Dir(path)
		if _, err := os.Stat(subdir); os.IsNotExist(err) {
			err := os.MkdirAll(subdir, 0755)
			if err != nil {
				return err
			}
		}

		if tmpl, ok := templates[tp]; ok {
			flags := os.O_RDWR | os.O_CREATE | os.O_TRUNC
			file, err := os.OpenFile(path, flags, 0664)
			if err != nil {
				return err
			}

			defer file.Close()

			fmt.Printf("Writing %s\n", fp)
			err = tmpl.Execute(file, vars)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
