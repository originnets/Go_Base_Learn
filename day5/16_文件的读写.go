package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

//写文件
func WriteFile(path string) {
	// 打开文件, 新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	// 使用完毕,需要关闭文件

	defer f.Close() //

	var buf string
	for i:=0; i<10;i++{
		// "i=1\n", 这个字符串存储再buf中
		buf = fmt.Sprintf( "i=%d\n", i)
		n, err1 := f.WriteString(buf)	//写入文件
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			return
		}
		fmt.Println("n = ", n)
	}
}

//读文件
func ReadFile(path string)  {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	buf := make([]byte, 1024*2) //2K大小

	//n代表从文件读取内容的长度

	n, err1 := f.Read(buf)
	if err1 != nil && err != io.EOF {	//文件出错, 同时没有到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))	//buf[:n]指定大小
}

//每次读取一行

func ReadFileLine(path string)  {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	//关闭文件
	defer f.Close()

	//新建一个缓冲区, 把内容先放在缓存区
	r := bufio.NewReader(f)

	for {
		//遇到\n结束读取 , 但是'\n'也读取进去
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF{ // 文件已经读完
				break
			}
			fmt.Println("err =", err)
		}
		// fmt.Printf("%s\n",string(buf))
		fmt.Printf("%s\n", strings.Replace( string(buf) , "\n", "", -1)) // 把\n过滤替换为空, string(buf), buf为切片类型
	}
}



func main()  {
	path := "./demo.txt"
	//WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)
}
