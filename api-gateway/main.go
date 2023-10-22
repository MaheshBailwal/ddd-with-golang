package main

import (
	reverseproxy "invafresh.com/api-gateway/reverse-proxy"
)

type requestInfo struct {
	path           string
	body           string
	verb           string
	contentLength  int
	command        string
	isRequestArray bool
}

func main() {
	reverseproxy.Start()
}
