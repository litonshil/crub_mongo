package main

import (
	container "crud_mongo/app"
	"crud_mongo/infra/config"
	"crud_mongo/infra/conn"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadConfig()
	conn.ConnectDb()
	e := echo.New()
	container.Init(e)
	e.Start(":8080")

}
