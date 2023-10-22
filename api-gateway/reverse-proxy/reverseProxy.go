package reverseproxy

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	routemap "invafresh.com/api-gateway/route-map"
)

type requestInfo struct {
	path           string
	body           string
	verb           string
	contentLength  int
	command        string
	isRequestArray bool
}

func Start() {

	logPath := "gatewaylogs.log"

	openLogFile(logPath)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	// initialize a reverse proxy and pass the actual backend server url here
	proxy, err := newProxy("http://localhost:8009")
	if err != nil {
		panic(err)
	}

	fmt.Println("listening on 4000")
	setUp()
	// handle all requests to your server using the proxy
	http.HandleFunc("/", ProxyRequestHandler(proxy))
	log.Fatal(http.ListenAndServe(":4000", nil))
}

func setUp() {
	routemap.Init()
	setIdentityProviderUrl()
}

// NewProxy takes target host and creates a reverse proxy
func newProxy(targetHost string) (*httputil.ReverseProxy, error) {
	url, err := url.Parse(targetHost)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(url)

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		modifyRequest(req)
	}

	proxy.ModifyResponse = modifyResponse()
	proxy.ErrorHandler = errorHandler()
	return proxy, nil
}

func modifyRequest(req *http.Request) requestInfo {

	req.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
	return reWriteRequest(req)
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
	return func(w http.ResponseWriter, req *http.Request, err error) {
		fmt.Printf("Got error while modifying response: %v \n", err)
		return
	}
}

// ProxyRequestHandler handles the http request using proxy
func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	}
}

func reWriteRequest(r *http.Request) requestInfo {
	fmt.Println(r.URL)
	if r.Body != nil {
		body, _ := ioutil.ReadAll(r.Body)

		err := r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		log.Println("OriginalRequest->", string(body))

		reqInfo := transfromRequestBody(string(body))
		r.Body = ioutil.NopCloser(strings.NewReader(reqInfo.body))
		urlPath := reqInfo.path
		r.Method = reqInfo.verb
		urlA, _ := url.Parse(urlPath)
		r.URL = urlA
		token := getToken(r.Header.Get("Cookie"))
		r.Header.Set("Authorization", token)
		r.Header.Set("command", reqInfo.command)

		r.ContentLength = int64(reqInfo.contentLength)
		fmt.Println("reqinfo->", reqInfo)
		log.Println("ModifiedRequest->", reqInfo.path)
		return reqInfo
	}

	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
	return requestInfo{}
}

func openLogFile(logfile string) {
	if logfile != "" {
		lf, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0640)

		if err != nil {
			log.Fatal("OpenLogfile: os.OpenFile:", err)
		}

		log.SetOutput(lf)
	}
}
