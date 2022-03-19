package main

import (
	"backend-go/db"
	"backend-go/handler"
	log "backend-go/log"
	"backend-go/repository/repo_impl"
	"backend-go/router"
	"github.com/labstack/echo/v4"
	"os"
)

func init(){
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}

func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "thinhnn",
		Password: "secret",
		DbName:   "testdb",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}



