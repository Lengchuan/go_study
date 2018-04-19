package main

import (
	"net/http"
	"io"
	"log"
	"os"
)

func main() {
	mux := http.NewServeMux()
	//自己实现一个handler注册到mux
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello2)

	//静态文件服务器

	wd, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/static/",
		//去除访问路径前缀						//文件前缀
		http.StripPrefix(("/static"), http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

//实现ServeHTTP()方法
func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL :"+r.URL.String())

}

func sayHello2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world 2")
}
