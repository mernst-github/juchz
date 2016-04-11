package main

import "net/http"

func init() {
  f := func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "/konzerte/nacht/index.html", http.StatusFound)
  }
  http.HandleFunc("/", f);
  http.HandleFunc("/index.html", f);
}

