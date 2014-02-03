package pages

import (
	"fmt"
	"foxsays/log"
	"os"
	"path/filepath"
	"strings"

	textTemplate "text/template"
)

type templateSet map[string]*textTemplate.Template

func (ts templateSet) Get(name string) Page {
	p := new(page)

	p.name = name
	p.template = ts[name]

	return p
}

func LoadPages(templatePath string) PageSet {
	t, e := loadTemplates(templatePath)
	log.FatalIf(e)
	return t
}

func loadTemplates(templatePath string) (ts templateSet, err error) {
	const pathSep = string(os.PathSeparator)

	ts = make(templateSet)

	if !strings.HasSuffix(templatePath, "pages") {
		err = fmt.Errorf("invalid template templatePath: %s", templatePath)
		return
	}

	var matches []string
	if matches, err = filepath.Glob(filepath.Join(templatePath, "*/main.html")); err != nil {
		return
	}

	for _, match := range matches {
		parts := strings.Split(match, pathSep)
		name := strings.Join(parts[len(parts)-3:len(parts)-1], pathSep)

		t := textTemplate.New(name)
		if _, err = t.ParseFiles(match); err != nil {
			return
		}
		ts[name] = t.Lookup("main.html")
	}

	return
}
