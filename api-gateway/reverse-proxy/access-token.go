package reverseproxy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var tokenMap map[string]string
var identityProviderUrl string

func getToken(cookie string) string {

	if tokenMap == nil {
		tokenMap = make(map[string]string)
	}

	sessionKey := findValue(cookie, "SessionKey")

	if len(tokenMap[sessionKey]) > 0 {
		fmt.Println("Token Found")
		return tokenMap[sessionKey]
	}
	userName := findValue(cookie, "Username")

	postBody, _ := json.Marshal(map[string]string{
		"userName": userName,
		"password": "abc",
	})
	responseBody := bytes.NewBuffer(postBody)
	log.Println("in to")
	resp, err := http.Post(identityProviderUrl, "application/json", responseBody)

	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	log.Println("in to2")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var bodyArr map[string]interface{}
	json.Unmarshal(body, &bodyArr)
	tokenMap[sessionKey] = bodyArr["Token"].(string)
	return tokenMap[sessionKey]
}

func findValue(cookie string, key string) string {

	arr := strings.Split(cookie, ";")
	key = fmt.Sprintf("%s=", key)
	for _, v := range arr {
		if strings.Contains(v, key) {
			return strings.Split(v, "=")[1]
		}
	}

	return ""
}

func setIdentityProviderUrl() {

	if len(identityProviderUrl) > 0 {
		return
	}

	content, _ := ioutil.ReadFile("service-config.json")
	var serviceConfig map[string]string
	json.Unmarshal(content, &serviceConfig)
	identityProviderUrl = serviceConfig["IDENTITY_PROVIDER"]
	value, found := os.LookupEnv("IDENTITY_PROVIDER")

	if found {
		identityProviderUrl = value
	}

	identityProviderUrl = fmt.Sprintf("http://%s:8080/api/v1/token", identityProviderUrl)
}
