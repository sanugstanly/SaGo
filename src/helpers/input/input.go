package input

import (
  "github.com/gorilla/mux"
  "net/http"
)

func GetParams(r *http.Request) map[string]string {
  param := mux.Vars(r)
  return param
}
