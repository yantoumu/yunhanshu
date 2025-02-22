  package main

  import (
      "io"
      "net/http"
  )

  func handler(w http.ResponseWriter, r *http.Request) {
      targetURL := r.URL.Query().Get("target")
      if targetURL == "" {
          http.Error(w, "Missing 'target' query parameter", http.StatusBadRequest)
          return
      }

      resp, err := http.Get(targetURL)
      if err != nil {
          http.Error(w, "Failed to reach target URL", http.StatusInternalServerError)
          return
      }
      defer resp.正文.Close()

      w.WriteHeader(resp.StatusCode)
      io.Copy(w, resp.Body)
  }

  func main() {
      http.HandleFunc("/", handler)
      http.ListenAndServe(":9000", nil)
  }
