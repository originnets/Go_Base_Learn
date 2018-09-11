package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)


type Cilent struct {
	C chan string	//用户发送数据的管道
	Name string		//用户名
	Addr string		//网络地址
}

//保存在线用户

//定义一个全局map
var onlineMap map[string]Cilent

//定义一个channel

var message = make(chan string)

func Manager() {
	// 给map分配空间
	onlineMap = make(map[string]Cilent)
	for {
		msg := <- message //没有消息前,这里回阻塞

		//遍历map,给map每个成员发送消息
		for  _, cli := range onlineMap {
			cli.C <- msg
		}

	}
}

func WriteMsgToClient(cli Cilent, conn net.Conn){
	for msg := range cli.C {
		conn.Write([]byte(msg))
	}
}

func MakeMsg(cli Cilent, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ":" + msg + "\n"
	return
}

func HandleConn(conn net.Conn) { //处理用户连接
	defer conn.Close()

	//获取客服端的网络地址

	cliAddr := conn.RemoteAddr().String()

	fmt.Println(conn.RemoteAddr().String(),"连接成功")
	//创建一个结构体,默认, 用户名和网络地址一样
	cli := Cilent{make(chan string), cliAddr, cliAddr}

	//把结构体添加到map中
	onlineMap[cliAddr] = cli

	//新开一个协程 ,专门给当前用户端发送消息
	go WriteMsgToClient(cli, conn)

	//广播某个人在线

	message <- MakeMsg(cli, "login")

	//提示,当前登入用户的身份,只给当前用户发送

	cli.C <- MakeMsg(cli, "I am here")

	isQuit := make(chan bool)	//用户是否主动退出
	hasDate := make(chan bool)	//用户是否有数据发送

	//新建一个协程, 接收用户发送过的数据
	go func() {
		buf := make([]byte, 1024*2)

		for {
			n, err := conn.Read(buf)
			if n == 0 { //对方断开,或者出问题
				isQuit <- true
				fmt.Println("conn.Read err = ", err)
				return
			}
			msg := string(buf[:n-1])

			if len(msg) == 3 && msg == "who" {
				//遍历map给当前用户发送所有成员
				conn.Write([]byte("user list:\n"))
				for _, tmp := range onlineMap {
					msg := tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			}else if len(msg) > 8 && msg[:6] == "rename"{
				//重命名
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			}else {
				//转发此内容
				message <- MakeMsg(cli, msg)
			}
			hasDate <- true	//代表有数据
		}
	}()
	for {
		//通过select检测channel的流动
		select {
		case <- isQuit:
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "logout\n")
			return
		case <-hasDate:	//有数据时走这个

		case <- time.After(60 * time.Second):	//60s后超时退出
			delete(onlineMap, cliAddr)
			message <- MakeMsg(cli, "timeout\n")
			return
		}
	}
}


func main() {
	//监听
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println("net.Listen err = ", err)
		return
	}

	defer listener.Close()

	//新开一个协程, 转发消息, 只要有消息来了 ,遍历map,给map每个成员发送消息
	go Manager()


	//主协程, 循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("net.Listen err = ", err)
			continue
		}
		go HandleConn(conn) // 处理用户连接
	}
}
