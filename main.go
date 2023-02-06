package main

import (
	"net/http"
	"text/template"
)

// Variável para carregar todos os templates que seja no formato HTML
var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	// Ouvir e responder no servidor
	http.ListenAndServe(":8000", nil)
}

// Função para
func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
