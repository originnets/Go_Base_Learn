package main

import (
	"fmt"
	"net/http"
)

//服务器编写的业务逻辑处理程序

func MyHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello World")
}


func main() {
	http.HandleFunc("/", MyHandler)
	//在制定的地址进行监听, 开启一个HTTP
	http.ListenAndServe("127.0.0.1:8000", nil)
}