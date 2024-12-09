package util

import (
	"ac9/glad/entity"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func GenerateTokens() (string, error) {
	// todo: move private variables to .env
	token_api := "https://aol-dev--awspoc.sandbox.my.salesforce.com/services/oauth2/token"
	client_id := "3MVG9u5bid8bKNSIXqYFxWiwRhWP07owBcYm4sK7E_I8J1R55euttZXX8PrDjbyI6qR1M8xqSfoeZKAnJrqZ4"
	client_secret := "C8BB85BA17737872D8B14322A85F297C1A30CA17ECC9253BECBE9BFDC1A192AF"

	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", client_id)
	data.Set("client_secret", client_secret)
	req, err := http.NewRequest("POST", token_api, bytes.NewBufferString(data.Encode()))
	if err != nil {
		log.Println("error creating the request", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") // this format encodes data as key-value pairs similar to query parameters in a url
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("an error occurred executing the request", err)
		return "", err
	}
	parse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("an error occurred when parsing")
		return "", err
	}
	log.Println("parse in gentoken", string(parse))
	var result entity.Token
	err = json.Unmarshal(parse, &result)
	if err != nil {
		log.Println("there was an error unmarshaling the json", err)
		return "", err
	}
	log.Println("token:", result.AuthToken)
	return result.AuthToken, nil
}
