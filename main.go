package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var logFile *os.File

func main() {

	r := gin.Default()

	gin.SetMode(gin.ReleaseMode)

	var err error

	if logFile, err = os.Create("logfile.txt"); err != nil {
		log.Fatal("Log file create:", err)
		return
	}

	r.Any("/*proxyPath", func(c *gin.Context) {

		for k, v := range c.Request.Header {
			stk := fmt.Sprintf("%v : %v \n", k, v)
			logFile.WriteString(stk)

		}

		c.JSON(http.StatusOK, gin.H{
			"X-Forwarded-For": c.Request.Header.Get("X-Forwarded-For"),
			"Remote Addr":     c.Request.RemoteAddr,
			"Path":            c.Request.URL.String(),
		})
	})

	r.Run()
}
