package handlers

import (
	"auth0Authentication/pkg/helpers"
	"auth0Authentication/pkg/models"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func RequestToken(c echo.Context) error {
	data := models.Login{}
	err := c.Bind(&data)
	responses := models.ApiResponse{}
	var payLoad = map[string]string{}

	if err != nil {
		c.JSON(http.StatusOK, models.ApiResponse{
			Status: http.StatusBadRequest,
		})
		return err
	}

	if os.Getenv("AUTHENTICATION_TYPE") == "OTP" {
		payLoad = map[string]string{
			"grant_type":    os.Getenv("PASSLSS_GRANT_TYPE"),
			"client_id":     os.Getenv("CLIENT_ID"),
			"client_secret": os.Getenv("CLIENT_SECRET"),
			"otp":           data.PasswordOrOTP,
			"realm":         os.Getenv("PASSLSS_REALM"),
			"username":      data.CustomerId,
			"scope":         os.Getenv("SCOPE"),
		}

	} else {
		payLoad = map[string]string{
			"grant_type":    os.Getenv("GRANT_TYPE"),
			"username":      data.CustomerId,
			"password":      data.PasswordOrOTP,
			"audience":      os.Getenv("AUDIENCE"),
			"client_id":     os.Getenv("CLIENT_ID"),
			"client_secret": os.Getenv("CLIENT_SECRET"),
			"realm":         os.Getenv("REALM"),
			"scope":         os.Getenv("SCOPE"),
		}
	}

	url := fmt.Sprintf("%v/oauth/token", os.Getenv("DOMAIN"))
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
