package main

import(
  "net/http"

  "github.com/gowebexamples/http-server/api"
) 


func main (){
  srv := api.NewServer()
  http.ListenAndServe(":8080", srv)
}

