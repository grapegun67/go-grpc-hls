package main

import (
	"fmt"
	"time"
	"net/url"
	"net/http"
	"net/http/httputil"
)

var (
	//http까지 붙여줘야함
	BACKEND_SERVER_LIST = []string{"http://127.0.0.1:8080", "http://127.0.0.1:8081"}
	BACKEND_SERVER_HANDLER []*httputil.ReverseProxy

	PROXY_SERVER_PORT	string = ":80"
)

/* ReverseProxy를 사용하기 위한 http의 인터페이스인 ServeHTTP를 정의	*/
type balanceHandler struct{}

func (b balanceHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	if time.Now().Unix() % 2 == 0 {
		BACKEND_SERVER_HANDLER[0].ServeHTTP(res, req)
	} else {
		BACKEND_SERVER_HANDLER[1].ServeHTTP(res, req)
	}

}

func main() {

	fmt.Println("[START]");

	/* 백엔드 서버 정의	*/
	for _, urls := range BACKEND_SERVER_LIST {
		backUrl, _ := url.Parse(urls)
		fmt.Println("Proxy server list:", backUrl)
		reverseHandler := httputil.NewSingleHostReverseProxy(backUrl)

		BACKEND_SERVER_HANDLER = append(BACKEND_SERVER_HANDLER, reverseHandler)
	}

	handler := balanceHandler{}

	/* Route to backendServer on LoadBalancing	*/
	http.ListenAndServe(PROXY_SERVER_PORT, handler)

	fmt.Println("[END]");
}
