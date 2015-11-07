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
    io.Copy(w, resp.Body)
  })
}
