package registration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func GetTokenFromRegistry(registryURL string) (interface{}, error) {
	// Create JSON request with username and password
	body := map[string]string{
		"username": "admin",
		"password": "admin",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// Get the token from the registry
	client := &http.Client{}
	req, err := client.Post(registryURL+"/login", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	respBody, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	// get the accessToken from the response
	var response map[string]interface{}
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return "", err
	}

	return response["accessToken"], nil

}

func RegisterSelf(registryURL string, accessToken string, appInfo map[string]interface{}) (interface{}, error) {

	// Create JSON request with username and password
	body := map[string]interface{}{
		"serviceName": appInfo["serviceName"],
		"port":        appInfo["port"],
		"description": appInfo["description"],
		"version":     appInfo["version"],
		"instanceId":  appInfo["instanceId"],
		"url":         appInfo["url"],
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}

	// Register the service with the registry
	// create request with headers
	client := &http.Client{}
	req, err := http.NewRequest("POST", registryURL+"/register", bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return nil, nil
}

func ReregisterSelf(registryURL string, accessToken string, appInfo map[string]interface{}) error {

	body := map[string]interface{}{
		"serviceName": appInfo["serviceName"],
		"instanceId":  appInfo["instanceId"],
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Register the service with the registry
	// create request with headers
	client := &http.Client{}
	req, err := http.NewRequest("POST", registryURL+"/reregister", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func DeregisterSelf(registryURL string, accessToken string, appInfo map[string]interface{}) error {
	body := map[string]interface{}{
		"serviceName": appInfo["serviceName"],
		"instanceId":  appInfo["instanceId"],
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// deregister the service with the registry
	// create request with headers
	client := &http.Client{}
	req, err := http.NewRequest("POST", registryURL+"/deregister", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// send request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
