package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// renderTemplate renders the template
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("error parsing %s template: %s \n", tmpl, err)
		return
	}
}
