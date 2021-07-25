package main

import (
	"html/template"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "Hello from Go server")
	renderTemplate(w, "index.html")

}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)
	log.Println("Server is running at 5000")
	http.ListenAndServe(":5000", nil)
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}
