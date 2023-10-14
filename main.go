package main

import (
	"fmt"
	"recycle/app/config"
	"recycle/app/database"
	"recycle/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	var cfg = config.InitConfig()

	dbMysql := database.InitMysqlConn(cfg)
	router.NewRoute(e, dbMysql)

	database.Migrate(dbMysql)
	
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVERPORT)))
}
