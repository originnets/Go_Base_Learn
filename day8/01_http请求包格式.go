package main

import (
	"fmt"
	"net"
)

func main()  {
	//监听
	listener ,err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}
	defer listener.Close()

	//阻塞等待用户连接

	conn, err1 := listener.Accept()
	if err1 != nil {
		fmt.Println("listener.Accept err1 = ", err1)
		return
	}
	defer conn.Close()

	//接收客服端的数据

	buf := make([]byte, 1024*4)
	n, err2 := conn.Read(buf)

	if n == 0 {
		fmt.Println("conn.Read err2 = ", err2)
		return
	}

	fmt.Printf("#%v#", string(buf[:n]))
}
