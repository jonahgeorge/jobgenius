package controllers

import "html/template"

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("views/_templates/*.html"))
	t = template.Must(t.ParseGlob("views/articles/*.html"))
	t = template.Must(t.ParseGlob("views/interviews/*.html"))
	t = template.Must(t.ParseGlob("views/interviews/components/*.html"))
	t = template.Must(t.ParseGlob("views/accounts/*.html"))
	t = template.Must(t.ParseGlob("views/users/*.html"))
	t = template.Must(t.ParseGlob("views/*.html"))
}
