package httprequest

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	response.Body.Close()

	return string(body), nil
}

func PostJson(url string, jsonData []byte) (string, error) {

	response, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	response.Body.Close()

	return string(body), nil
}

// parseJSONResponse parses a string http response to json format
func parseJsonResponse(res string, object *interface{}) error {

	resBytes := []byte(res)
	// declaring a map for key names as string and values as interface{}
	var jsonResponse map[string]interface{}
	err := json.Unmarshal(resBytes, &jsonResponse)

	if err != nil {
		return err
	}

	return nil
}
