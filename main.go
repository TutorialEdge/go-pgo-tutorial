package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type HomePage struct {
	Title   string
	Content string
}

func main() {
	fmt.Println("Go PGO Tutorial")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		tmpl.Execute(w, HomePage{
			Title:   "My Awesome Website",
			Content: "All my awesome content",
		})
	})
	http.ListenAndServe(":9000", mux)
}
