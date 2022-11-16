package server

import (
    "fmt"
    "net/http"
)

func Serve() {
  err := http.ListenAndServe(":8080", nil)
  if err != nil {
    fmt.Println( "there was an error" )
  }
}
