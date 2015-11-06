package main

import "net/http"

func init() {
  http.HandleFunc("/calendar/", func (w http.ResponseWriter, r *http.Request) {
    u := *r.URL
    u.Scheme = "https"
    u.Host = "www.google.com"
    http.Redirect(w, r, u.String(), http.StatusFound)
  })
}
