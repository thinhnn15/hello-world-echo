package main

import (
	"backend-go/db"
	"backend-go/handler"
	"context"
	"github.com/labstack/echo/v4"
	log "backend-go/log"
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

	var email string
	err := sql.Db.GetContext(context.Background(), &email, "SELECT email FROM users WHERE email=$1", "abc@gmail.com")
	if err != nil{
		log.Error(err.Error())
	}
	
	e := echo.New()
	e.GET("/", handler.Welcome)
	e.GET("/user/sign-in", handler.HandleSignIn)
	e.GET("/user/sign-up", handler.HandleSignUp)
	e.Logger.Fatal(e.Start(":3000"))
}



