package main

import "net/http"

func init() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/konzerte/warum/flyer.html", http.StatusFound)
  })
}
