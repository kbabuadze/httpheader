package main

import (
	"fmt"
	"log"
	"net/http"
)

// LogInfo Logs IP and Header to the console  and Displays it on a website
func LogInfo(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		fmt.Printf("%v : %v \n", k, v)
	}

	fmt.Println("IP Address received by the Server: " + r.RemoteAddr)

	w.WriteHeader(200)
	fmt.Fprintln(w, "X-Forwarded-For received by the Server: "+r.Header.Get("X-Forwarded-For"))
	fmt.Fprintln(w, "IP Address received by the Server: "+r.RemoteAddr)

}

func main() {

	http.HandleFunc("/", LogInfo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
