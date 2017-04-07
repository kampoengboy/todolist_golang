package main

import "html/template"

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))
