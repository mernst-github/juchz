package main

import ("io"
        "net/http"
        "appengine"
        "appengine/urlfetch")

func init() {
  http.HandleFunc("/calendar/", func (w http.ResponseWriter, r *http.Request) {
    u := *r.URL
    u.Scheme = "https"
    u.Host = "www.google.com"
    
    c := appengine.NewContext(r)
    client := urlfetch.Client(c)
    resp, err := client.Get(u.String())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()
    w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
    w.Header().Set("Expires", resp.Header.Get("Expires"))
    w.Header().Set("Cache-Control", resp.Header.Get("Cache-Control"))
    w.Header().Set("Pragma", "no-cache")
    io.Copy(w, resp.Body)
  })
}
