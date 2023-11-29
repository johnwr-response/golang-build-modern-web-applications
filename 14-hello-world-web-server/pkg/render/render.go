package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderingTemplate renders templates using html/template
func RenderingTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Printf("error parsing template: %s \n", err)
		//fmt.Printf("error parsing %s template: %s \n", tmpl, err)
		return
	}
}
