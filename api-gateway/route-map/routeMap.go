package routemap

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type RequestRouteInfo struct {
	Url            string
	Verb           string
	IsRequestArray bool
}

var requestRouteMap map[string]RequestRouteInfo

func GetRoute(command string) RequestRouteInfo {
	return requestRouteMap[command]
}

func Init() {

	content, _ := ioutil.ReadFile("route-map/routeMap.json")
	var routeMap map[string]string
	json.Unmarshal(content, &routeMap)

	content, _ = ioutil.ReadFile("service-config.json")
	var serviceConfig map[string]string
	json.Unmarshal(content, &serviceConfig)
	fmt.Println("sc->", serviceConfig)

	requestRouteMap = make(map[string]RequestRouteInfo)
	for k := range routeMap {
		arr := strings.Split(routeMap[k], ",")
		routeInfo := RequestRouteInfo{}
		routeInfo.Url = replaceServicePlaceHolder(arr[0], serviceConfig)
		routeInfo.Verb = arr[1]
		routeInfo.IsRequestArray = arr[2] == "RQUEST_ARRAY"
		requestRouteMap[k] = routeInfo
	}

	fmt.Println("rm->", requestRouteMap)
}

func replaceServicePlaceHolder(url string, serviceConfig map[string]string) string {
	for k := range serviceConfig {
		servicePlaceHolder := fmt.Sprintf("{%s}", k)
		serviceUrl := serviceConfig[k]
		value, found := os.LookupEnv(k)

		if found {
			serviceUrl = value
		}

		url = strings.ReplaceAll(url, servicePlaceHolder, serviceUrl)
	}

	return url
}
