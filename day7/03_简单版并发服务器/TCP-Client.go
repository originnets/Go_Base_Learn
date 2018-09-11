package main

import (
	"fmt"
	"net"
	"os"
)

func main()  {
	//主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("Dial_err = ", err)
		return
	}

	//main调用完毕,关闭连接
	defer conn.Close()

	go func() {

		//从键盘输入内容,给服务发送内容
		str := make([]byte, 2048)
		for {
			n, err := os.Stdin.Read(str)	//从键盘读取内容
			if err != nil {
				fmt.Println("Stdin.Read_err = ", err)
				return
			}
			//fmt.Println(string(str[:n])) //打印接收到的内容
			//把输入的内容发送给服务器
			conn.Write([]byte(str[:n]))
		}
	}()

	//接收服务器回复的数据,新任务
	//切片缓存
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)	//接收服务器的请求
		if err != nil {
			fmt.Println("Read_err = ", err)
			return
		}
		fmt.Println(string(buf[:n])) //打印接收到的内容, 转换为字符串
	}
}