package handlers

import (
	"auth0Authentication/pkg/helpers"
	"auth0Authentication/pkg/models"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Signup(c echo.Context) error {
	data := models.Signup{}
	responses := models.ApiResponse{}
	err := c.Bind(&data)

	if err != nil {
		c.JSON(http.StatusOK, models.ApiResponse{
			Status: http.StatusBadRequest,
		})
		return err
	}

	payLoad := map[string]string{
		"username":   data.UserName,
		"email":      data.Email,
		"password":   data.Password,
		"connection": os.Getenv("REALM"),
	}
	url := fmt.Sprintf("%v/api/v2/users", os.Getenv("DOMAIN"))
	accesToken := helpers.ManagementTokenGenerator()
	header := map[string]string{"Authorization": fmt.Sprintf("Bearer %v", accesToken), "Content-Type": "application/json"}
	response := helpers.HttpRequest(url, http.MethodPost, payLoad, header)

	if response["status"] == 201 {
		responses.Status = response["status"].(int)
		responses.Error = ""
		responses.Data = response["data"]
		c.JSON(http.StatusOK, response)
	} else {
		responses.Status = response["status"].(int)
		responses.Error = fmt.Sprintf("%v", response["error"])
		c.JSON(http.StatusNotAcceptable, response)
	}

	return nil
}
