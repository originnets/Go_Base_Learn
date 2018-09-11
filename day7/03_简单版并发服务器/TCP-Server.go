package main

import (
	"fmt"
	"net"
	"strings"
)

// 处理用户请求

func HandleConn(conn net.Conn) {

	//获取客服端的网络地址信息
	addr := conn.RemoteAddr().String()
	fmt.Println("连接成功 : ", addr)
	conn.Write([]byte(string("连接成功,请输入:")))

	defer conn.Close()

	//循环读取用户数据
	for {
	buf := make([]byte, 2048)

	n, err := conn.Read(buf)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	fmt.Printf("[%s]:[%s]\n", addr, string(buf[:n]))

	if string(buf[:n]) == "exit\r\n" { //windows下多了两个字符\r\n
 		fmt.Printf("%s--用户退出", addr)
		return
	}
	// 把数据转换为一个大写,在发送给用户
	conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}
}


func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer  listener.Close()

	//接收多个用户

	for {
		conn, err1 := listener.Accept()
		if err != nil {
			fmt.Println("err1 = ", err1)
			return
		}

		//处理用户的请求, 新建一个协程
		go HandleConn(conn)

	}

}
