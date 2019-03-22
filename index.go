package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logFile *os.File

// LogInfo Logs IP and Header to the console  and Displays it on a website
func LogInfo(w http.ResponseWriter, r *http.Request) {

	for k, v := range r.Header {
		stk := fmt.Sprintf("%v : %v \n", k, v)
		logFile.WriteString(stk)

	}

	logFile.WriteString("================== \n")

	fmt.Println("IP Address received by the Server: " + r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintln(w, "X-Forwarded-For received by the Server: "+r.Header.Get("X-Forwarded-For"))
	fmt.Fprintln(w, "IP Address received by the Server: "+r.RemoteAddr)

}

func main() {
	var err error
	logFile, err = os.Create("logfile.txt")
	if err != nil {
		log.Fatal("Log file create:", err)
		return
	}

	http.HandleFunc("/", LogInfo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
