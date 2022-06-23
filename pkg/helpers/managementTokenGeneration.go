package helpers

import (
	"fmt"
	"net/http"
	"os"
)

func ManagementTokenGenerator() string {
	payLoad := map[string]string{
		"grant_type":    os.Getenv("ADM_GRANT_TYPE"),
		"audience":      os.Getenv("AUDIENCE"),
		"client_id":     os.Getenv("ADM_CLIENT_ID"),
		"client_secret": os.Getenv("ADM_CLIENT_SECRET"),
	}

	url := fmt.Sprintf("%v/oauth/token", os.Getenv("DOMAIN"))
	header := map[string]string{"Content-Type": "application/json"}
	response := HttpRequest(url, http.MethodPost, payLoad, header)
	responseData := response["data"].(map[string]interface{})
	if responseData["error"] != nil {
		return responseData["error"].(string)
	} else {
		return responseData["access_token"].(string)
	}

}
