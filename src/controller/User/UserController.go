package userController

import (
  "strconv"
  "net/http"
  "model/User"
  "interface/User"
)

var obj = []userInterface.User{userModel.Details{"Your Name", 25}}

func GetName(w http.ResponseWriter, r *http.Request)  {
  name := obj[0].PrintUserName()
  w.Write([]byte(name))
}

func GetAge(w http.ResponseWriter, r *http.Request)  {
  age := obj[0].PrintUserAge()
  w.Write([]byte(strconv.Itoa(age)))
}
