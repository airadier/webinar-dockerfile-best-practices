package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var logFile *os.File

func main() {
	var err error
	logFile, err = os.OpenFile("access.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer logFile.Close()

	http.HandleFunc("/", logAndGreet)
	fmt.Println("Listening at :5000")
	if err = http.ListenAndServe(":5000", http.DefaultServeMux); err != nil {
		panic(err)
	}
}

func logAndGreet(w http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	logFile.WriteString(fmt.Sprintf("[%s] IP=%s name=%s\n", time.Now().Format(time.RFC3339), request.RemoteAddr, request.Form.Get("name")))
	if request.Form.Get("name") != "" {
		w.Write([]byte(fmt.Sprintf("Hi %s!\n", request.Form.Get("name"))))
	} else {
		w.Write([]byte("Hi anonymous!\n"))
	}
}