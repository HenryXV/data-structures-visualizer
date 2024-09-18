package main

import (
	"data_structures_visualizer/pkg/lists/linked_list"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Viz struct {
	Image64 string
}

func main() {
	var viz Viz

	ll := linked_list.New()
	tmpl := template.Must(template.ParseFiles("pkg/ui/index.html", "pkg/ui/image.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, viz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		ll.Add(r.FormValue("data"))
		buf := ll.Visualizer()

		image64 := base64.StdEncoding.EncodeToString(buf.Bytes())

		err := tmpl.ExecuteTemplate(w, "image.html", image64)
		if err != nil {
			log.Fatal(err)
		}
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
