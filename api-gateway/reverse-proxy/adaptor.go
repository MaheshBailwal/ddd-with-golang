package reverseproxy

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	routemap "invafresh.com/api-gateway/route-map"
)

var targetBodyMap map[string][]byte

func modifyResponseFile() func(*http.Response) error {
	return func(resp *http.Response) error {

		content, _ := ioutil.ReadFile("mapping/test12.json")
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		resp.Body = ioutil.NopCloser(strings.NewReader(string(content)))
		resp.Header.Set("Content-Length", strconv.Itoa(len(content)))
		resp.Header.Set("Content-Type", "application/json")
		return nil
	}
}

func transfromRequestBody(body string) requestInfo {
	reqInfo := requestInfo{}
	var bodyArray map[string]string
	json.Unmarshal([]byte(body), &bodyArray)
	command := ""

	if len(bodyArray["msg"]) > 0 {
		command = fmt.Sprintf("%s_%s", bodyArray["msg"], bodyArray["cco"])
	} else {
		command = fmt.Sprintf("%s_%s", bodyArray["MSG"], bodyArray["CCO"])
	}
	reqInfo.command = command
	command = fmt.Sprintf("%s_REQUEST", command)

	fmt.Println("comand->", command)

	// content, _ := ioutil.ReadFile("mapping/routeMap.json")
	// var routeMap map[string]string
	// json.Unmarshal(content, &routeMap)
	// arr := strings.Split(routeMap[command], ",")
	// reqInfo.path = arr[0]
	// reqInfo.verb = arr[1]
	// reqInfo.isRequestArray = arr[2] == "RQUEST_ARRAY"
	routeInfo := routemap.GetRoute(command)
	fmt.Println("route->", routeInfo)
	reqInfo.path = routeInfo.Url
	reqInfo.verb = routeInfo.Verb
	reqInfo.isRequestArray = routeInfo.IsRequestArray

	targetBody := getTargetBody(command)
	if len(targetBody) < 2 {
		return reqInfo
	}

	var bodyArr map[string]interface{}
	json.Unmarshal([]byte(body), &bodyArr)
	payload := ""
	if bodyArr["PostArray"] != nil {
		items := bodyArr["PostArray"].([]interface{})
		payload = transformArray(items, body, targetBody, reqInfo.isRequestArray)
	} else {

		payload = replacePlaceHolders(body, targetBody, bodyArr)
	}

	reqInfo.body = payload
	reqInfo.contentLength = len(payload)
	return reqInfo
}

func modifyResponse() func(*http.Response) error {
	return func(resp *http.Response) error {
		command := resp.Request.Header.Get("command")
		body, _ := ioutil.ReadAll(resp.Body)
		reqInfo := transfromResponseBody(string(body), command)
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		resp.Body = ioutil.NopCloser(strings.NewReader(reqInfo.body))
		resp.Header.Set("Content-Length", strconv.Itoa(reqInfo.contentLength))
		resp.Header.Set("Content-Type", "application/json")

		log.Println("OriginalResponse->", string(body))
		log.Println("ModifiedResponse->", reqInfo.body)

		return nil
	}
}

func transfromResponseBody(body string, command string) requestInfo {
	reqInfo := requestInfo{}
	var bodyArray map[string]string
	json.Unmarshal([]byte(body), &bodyArray)
	command = fmt.Sprintf("%s_RESPONSE", command)

	fmt.Println("command->", command)

	targetBody := getTargetBody(command)

	if len(targetBody) < 2 {
		return reqInfo
	}

	var bodyArr map[string]interface{}
	json.Unmarshal([]byte(body), &bodyArr)

	originalPayload := bodyArr["payload"].([]interface{})

	payload := transformArray(originalPayload, body, targetBody, true)

	reqInfo.body = payload
	reqInfo.contentLength = len(payload)
	return reqInfo
}

func getTargetBody(command string) []byte {

	if targetBodyMap == nil {
		targetBodyMap = make(map[string][]byte)
	}

	val, ok := targetBodyMap[command]

	if ok {
		return val
	}

	folder := strings.Split(command, "_")[0]
	fileName := fmt.Sprintf("mapping/%s/%s.json", folder, command)
	targetBody, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	targetBodyMap[command] = targetBody
	return targetBody
}

func transformArray(items []interface{}, body string, targetBody []byte, isArray bool) string {
	response := ""
	for _, item := range items {
		mapItem := item.(map[string]interface{})
		itemResponse := replacePlaceHolders(body, targetBody, mapItem)
		response = fmt.Sprintf("%s \n %s,", response, itemResponse)
	}
	response = strings.Trim(response, ",")
	if isArray {
		response = fmt.Sprintf("[%s]", response)
	}
	return response
}
func replacePlaceHolders(text string, content []byte, postArray map[string]interface{}) string {
	text = string(content)

	for k := range postArray {
		prop := fmt.Sprintf("{%s#}", k)
		val := fmt.Sprintf("%v", postArray[k])
		text = strings.ReplaceAll(text, prop, val)
	}

	scanner := bufio.NewScanner(strings.NewReader(text))
	var playload = ""
	for scanner.Scan() {
		var line = scanner.Text()

		if !strings.Contains(line, "#}") {
			playload = fmt.Sprintf("%s\n %s", playload, line)
		}
	}
	return playload
}
