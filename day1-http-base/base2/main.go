package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine is a universal handler for all request
type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		// %q 会返回带有引号的字符串输出
		fmt.Fprintf(w, "req.URL.Path is %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "header[%q] = [%q]\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 Not Found :%s", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9989", engine))
}
