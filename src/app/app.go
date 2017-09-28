package app

import (
  "routes"
  "net/http"
  "fmt"
)

func Start()  {
  router := routes.RenderRoutes()
  err := http.ListenAndServe(":9081", router)
  fmt.Print(err)
}
