package main

import (
	"fmt"
	"net/http"
)

//w , 给歌赋端回复数据
//r, 读取客服端发送的数据

func HandConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method = ", r.Method)
	fmt.Println("r.URL = ", r.URL)
	fmt.Println("r.Header = ", r.Header)
	fmt.Println("r.Body = ", r.Body)
	fmt.Println("r.RemoteAddr = ", r.RemoteAddr)
	w.Write([]byte("Hello Go")) //给客服端恢复数据
}

func main()  {
	//注册处理函数, 用户连接,自动调用指定的处理函数
	http.HandleFunc("/", HandConn)
	//监听绑定
	http.ListenAndServe(":8000", nil)
}
