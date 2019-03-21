package main

import (
	"fmt"
	"log"
	"net/http"
)

// LofInfo Logs IP and Header to the console
// and Displays it on a website
func LogInfo(w http.ResponseWriter, r *http.Request) {

	fmt.Println("\n\nX-Forwarded-For received by the Server: " + r.Header.Get("X-Forwarded-For"))
	fmt.Println("IP Address received by the Server: " + r.RemoteAddr)

	w.WriteHeader(200)
	fmt.Fprintln(w, "X-Forwarded-For received by the Server: "+r.Header.Get("X-Forwarded-For"))
	fmt.Fprintln(w, "IP Address received by the Server: "+r.RemoteAddr)

}

func main() {

	http.HandleFunc("/", LogInfo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
