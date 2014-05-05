package page

import (
	"html/template"
	"io"
	"foxsays/config"
	"foxsays/log"
)

type Page struct {
	Id   string
	Head template.HTML
	Main template.HTML
	Tail template.HTML
}

func (p Page) CSS() []string {
	if p.Id == "" {
		return []string{}
	}

	return config.Assets.Get(p.Id + "/main.min.css")
}

func (p Page) JS() []string {
	if p.Id == "" {
		return []string{}
	}

	return config.Assets.Get("vendor-main.min.js", p.Id+"/main.min.js")
}

func (p Page) Render(w io.Writer) {
	render(w, p)
}

func render(w io.Writer, p Page) {
	log.PanicIff(layout.Execute(w, p), "page.render failed for page.Id %q", p.Id)
}

var layout *template.Template

func init() {
	tmpl, err := template.New("layout").Parse(layoutHTML)
	log.PanicIf(err)
	layout = tmpl
}

const layoutHTML = `<!DOCTYPE html><html lang="en"><head>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1">
{{range .CSS}}<link rel="stylesheet" href="{{.}}">{{end}}
{{.Head}}
</head><body>
<div id="main">{{.Main}}</div>
{{.Tail}}
<script src="//ajax.googleapis.com/ajax/libs/jquery/2.1.0/jquery.min.js"></script>
<script src="//netdna.bootstrapcdn.com/bootstrap/3.1.1/js/bootstrap.min.js"></script>
{{range .JS}}<script src="{{.}}"></script>{{end}}
</body></html>`
