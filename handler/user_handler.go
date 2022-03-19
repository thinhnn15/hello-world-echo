package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HandleSignIn(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"user":  "ThinhNN",
		"email": "thinhnnfs5@gmail.com",
	})
}
// User
type User struct {
	Name  string `json:"name" xml:"name"`
	Email string `json:"email" xml:"email"`
}
func HandleSignUp(c echo.Context) error {
	u := User{
		Name:  "ThinhNN",
		Email: "thinhnnfs5@gmail.com",
	}
	return c.JSON(http.StatusOK, u)
}
