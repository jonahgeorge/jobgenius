package controllers

import "html/template"

var t *template.Template

func init() {
	t = template.Must(t.ParseGlob("templates/shared/*.html"))
	t = template.Must(t.ParseGlob("templates/articles/*.html"))
	t = template.Must(t.ParseGlob("templates/interviews/*.html"))
    t = template.Must(t.ParseGlob("templates/users/*.html"))
	t = template.Must(t.ParseGlob("templates/static/*.html"))
}
