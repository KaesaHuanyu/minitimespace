package handlers

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//Login create a token for current user
func (h *Handler) Login(c echo.Context) (err error) {
	username := c.QueryParam("username")
	//jwt
	//create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": username,
		"userID":   "8823",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	//encode signed
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return
	}

	return c.JSONPretty(http.StatusOK, map[string]string{
		"token": t,
	}, "	")
}

//Restricted is
func (h *Handler) Restricted(c echo.Context) (err error) {
	userToken := c.Get("user").(*jwt.Token)
	userClaims := userToken.Claims.(jwt.MapClaims)
	username := userClaims["userName"].(string)
	userID := userClaims["userID"].(string)
	return c.String(http.StatusOK, "Welcome "+username+`!
		Your ID is `+userID)
}

//Protect is used to protect data
func (h *Handler) Protect(handlerFunc echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		//check if user is logged in
		//TODO
		return handlerFunc(c)
	}
}
