package handlers

import (
	"auth0Authentication/pkg/helpers"
	"auth0Authentication/pkg/models"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func RequestOTP(c echo.Context) error {
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
		"client_id":     os.Getenv("CLIENT_ID"),
		"client_secret": os.Getenv("CLIENT_SECRET"),
		"connection":    os.Getenv("CONNECTION"),
		"send":          os.Getenv("SEND"),
		"email":         data.Email,
	}
	url := fmt.Sprintf("%v/passwordless/start", os.Getenv("DOMAIN"))
	header := map[string]string{"Content-Type": "application/json"}
	response := helpers.HttpRequest(url, http.MethodPost, payLoad, header)

	if response["status"] == 200 {
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
