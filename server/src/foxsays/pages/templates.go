package pages

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)

type templateSet map[string]*template.Template

var templates templateSet

func LoadTemplates(path string) (err error) {
	templates, err = loadTemplates(path)
	return
}

func loadTemplates(path string) (ts templateSet, err error) {
	const pathSep = string(os.PathSeparator)

	ts = make(templateSet)

	if !strings.HasSuffix(path, "pages") {
		err = fmt.Errorf("invalid template path: %s", path)
		return
	}

	var matches []string
	if matches, err = filepath.Glob(filepath.Join(path, "*/main.html")); err != nil {
		return
	}

	for _, match := range matches {
		parts := strings.Split(match, pathSep)
		name := strings.Join(parts[len(parts)-3:len(parts)-1], pathSep)

		t := template.New(name)
		if _, err = t.ParseFiles(match); err != nil {
			return
		}
		ts[name] = t.Lookup("main.html")
	}

	return
}
