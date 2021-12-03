package main

import (
	"fmt"
	"net/http"
	"log"
)

var (
        BACKEND_SERVER_IP       string = "127.0.0.1:8081"
)

func main() {

	fmt.Println("[START]");
	fmt.Println("Server ip:", BACKEND_SERVER_IP);

	editb := fmt.Sprintf("<%v> server called by the Reverse Proxy", BACKEND_SERVER_IP)

	/* 백엔드 서버 핸들러 및 속성 정의	*/
	backendServer := &http.Server{
		Addr:		BACKEND_SERVER_IP,
		Handler:	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					fmt.Fprintln(w, editb)
				}),
	}
	defer backendServer.Close()

	log.Fatal(backendServer.ListenAndServe())

	fmt.Println("[END]");
}
