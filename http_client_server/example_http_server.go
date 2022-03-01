package httpclientserver

import (
	"fmt"
	"net/http"
)

func InitHttpServer() {
	fmt.Println("Strting Web Server")
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello Golang Server"))
	})

	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":8090", nil)
}

func Hello(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte(`{"id":1,"name":"Ram"}`))
}
