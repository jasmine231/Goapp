package main

import {
    "fmt"
    "os"
    "log"
    "net/http"
}

func handler(w http.ResponseWriter, r *http.request){
  name, err := os.Hostname()
  if err != nil {
    panic(err)
  }
  
  fmt.Fprintln(w, "hostname:", name)
}

func main() {
  fmt.Fprintln(os.Stdout, "Starting GoApp Server...")
  http.HanlderFunc("/",handler)
  log.fatal(http.ListenAndServe(":8080",nil))
}
