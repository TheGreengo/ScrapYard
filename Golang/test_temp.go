package main

import (
    "html/template"
    "net/http"
)

type Page struct {
    Title string
    Text string
}

func main() {
    pg := Page{Title:"The Name", Text:"was successful"}
    tmp := template.Must(template.ParseFiles("test.html"))
    
    http.HandleFunc("/test", func (w http.ResponseWriter, r *http.Request) {
        tmp.Execute(w, pg)
    })

    http.ListenAndServe(":420", nil)
}
