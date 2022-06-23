package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpRequest(url string, method string, inputData map[string]string, header map[string]string) map[string]interface{} {
	var token interface{}

	jsonStr, err := json.Marshal(inputData)

	if err != nil {
		return map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err,
			"data":   "",
		}
	}

	// payload := strings.NewReader(inputData)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

	if err != nil {
		return map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err,
			"data":   "",
		}
	}

	for key, val := range header {
		req.Header.Add(key, val)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err,
			"data":   "",
		}
	}
	fmt.Println("Response Status = ", res.StatusCode)

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err,
			"data":   "",
		}
	}

	err = json.Unmarshal([]byte(body), &token)

	if err != nil {
		return map[string]interface{}{
			"status": http.StatusBadRequest,
			"error":  err,
			"data":   "",
		}
	}

	return map[string]interface{}{
		"status": res.StatusCode,
		"error":  nil,
		"data":   token,
	}
}
