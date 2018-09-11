package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

//接收文件内容
func RecvFile(FileName string, conn net.Conn) {
	// 新建文件
	f, err := os.Create(FileName)
	if err != nil {
		fmt.Println("os.Create err= ", err)
		return
	}

	defer f.Close()

	buf := make([]byte, 1024*4)

	//接收文件,写多少,接收多少

	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println(FileName ,"文件发送完毕")
				return
			} else {
				fmt.Println("conn.Read err1 = ", err1)
				return
			}
		}
		f.Write(buf[:n]) //往文件写入内容
	}

}

func main() {

	//设置监听
	lister, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err= ", err)
		return
	}

	defer lister.Close()

	//阻塞等待用户连接
	for {
		conn, err1 := lister.Accept()
		if err1 != nil {
			fmt.Println("lister.Accept err1= ", err1)
			return
		}

		defer conn.Close()

		buf := make([]byte, 1024)

		n, err2 := conn.Read(buf)	//读取对方发送的文件名
		if err2 != nil {
			fmt.Println("lister.Accept err2= ", err2)
			return
		}
		FileName := string(buf[:n])	//接收文件名
		//回复"ok"
		conn.Write([]byte("ok"))

		//接收文件内容
		go RecvFile(FileName, conn)
	}
}
