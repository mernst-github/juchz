package main

import "net/http"

func init() {
  http.HandleFunc("/intranet", func (w http.ResponseWriter, r *http.Request) {
    http.Redirect(w, r, "intranet.html", http.StatusMovedPermanently)
  })
}


