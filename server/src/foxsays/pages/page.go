package pages

import (
	"bytes"
	"encoding/json"
	"fmt"
	"foxsays/log"
	"html/template"
	"io"
)

type (
	Page interface {
		WriteTitle(t string)
		WriteHead(string)
		WriteMain(string)
		WriteTail(string)
		WritePool(id string, pool interface{})
		Render(io.Writer)
	}

	page struct {
		name     string
		Head     bytes.Buffer
		Main     bytes.Buffer
		Tail     bytes.Buffer
		template *template.Template
	}
)

func GetPage(name string) Page {
	p := new(page)

	p.name = name
	p.template = templates[name]

	return p
}

func (p *page) WriteTitle(title string) {
	p.Head.WriteString(`<title>`)
	p.Head.WriteString(title)
	p.Head.WriteString(`</title>`)
}

func (p *page) WriteHead(s string) { p.Head.WriteString(s) }
func (p *page) WriteMain(s string) { p.Main.WriteString(s) }
func (p *page) WriteTail(s string) { p.Tail.WriteString(s) }

func (p *page) WritePool(id string, pool interface{}) {
	p.Tail.WriteString(fmt.Sprintf(`<script id=%q type="application/json">`, id))

	b, e := json.Marshal(pool)
	log.PanicIff(e, "page.WritePool failed %v", p.name)

	p.Tail.Write(b)
	p.Tail.WriteString(`</script>`)
}

func (p *page) Render(w io.Writer) {
	if p.template == nil {
		panic(fmt.Errorf("missing template for %v", p.name))
	}

	e := p.template.Execute(w, struct {
		Head template.HTML
		Main template.HTML
		Tail template.HTML
	}{
		Head: template.HTML(p.Head.String()),
		Main: template.HTML(p.Main.String()),
		Tail: template.HTML(p.Tail.String()),
	})

	log.PanicIff(e, "page.Render failed %v", p.name)
}
