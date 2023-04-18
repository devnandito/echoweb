package main

import (
	"github.com/devnandito/echoweb/handlers"
	"github.com/labstack/echo"
)

func main(){
	e := echo.New()
	e.GET("/", handlers.HandlerShowModule)
	e.Logger.Fatal(e.Start(":8080"))
}