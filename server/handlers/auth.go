package handlers

import (
	"saltcrypt/database"
	"saltcrypt/types"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)



func RegisterHandler(c echo.Context) error {
	
	u := new(types.User)
	if err := c.Bind(u); err != nil {
	  return c.String(http.StatusBadRequest, "bad request")
	}

	user := types.UserDTO{
		Email: u.Email,
		Password: u.Password,
	}

	if err := database.RegisterDB(user); err != nil{
		log.Printf("Error Registering User: %v", err)
		return err
	}

	return c.JSON(http.StatusCreated, "User Registered Successfully")
}


func LoginHandler(c echo.Context) error {

	u := new(types.User)
	if err := c.Bind(u); err != nil {
	  return c.String(http.StatusBadRequest, "bad request")
	}

	user := types.UserDTO{
		Email: u.Email,
		Password: u.Password,
	}

	fmt.Printf("Email sa handler: %s", user.Email)
	fmt.Printf("Password sa handler: %s", user.Password)

	if err := database.LoginDB(user); err != nil{

		if err.Error() == "invalid_credentials"{
			return c.JSON(http.StatusNotFound, "The credentials you entered are incorrect. Please check your details and try again.")
		}

		return c.JSON(http.StatusInternalServerError, "Internal Server Error")

	}

	return c.JSON(http.StatusOK, user.Email)
}


func FetchUsersHandler(c echo.Context) error {
	users, err := database.FetchUsers()
	if err != nil {
		log.Printf("Error Fetching User: %v", err)
		return c.JSON(http.StatusInternalServerError, "Error Fetching Users")
	}
	return c.JSON(http.StatusOK, users)
}