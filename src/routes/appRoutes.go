package routes

import (
  "controller/Hello"
  "controller/User"
  "model"
)

var routes = routesModel.Routes{
    routesModel.Route{"HelloWorld", "GET", "/HelloWorld", helloController.HelloWorld,},
    routesModel.Route{"GetName", "GET", "/user/name", userController.GetName,},
    routesModel.Route{"GetAge", "GET", "/user/age", userController.GetAge,},
}
