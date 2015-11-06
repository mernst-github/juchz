package main

import "net/http"

func init() {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "https://www.google.com" + r.Url.EscapedPath(), http.StatusFound)
  })
}
