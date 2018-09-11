package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func HttpGet(url string)(result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页body内容
	buf := make([]byte, 1024*4)
	for {
		n , err := resp.Body.Read(buf)
		if n == 0 {	//读取结束, 或者出问题了
			if err == io.EOF {
				//fmt.Println("爬取完成")
			}else {
				fmt.Println("resp.Body.Read err = ", err)
			}
			break
		}
		result += string(buf[:n])
	}
	return
}


//爬取一个网页

func SpiderPape(i int, page chan int){
	url := "https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
	fmt.Println("url = ", url)
	fmt.Printf("正在爬第%d页网页:%s\n", i, url)

	//爬(将所有的网站的内容全部爬下来)
	result ,err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	//把内容写到文件
	FileName := strconv.Itoa(i) + ".html"
	f, err1 := os.Create(FileName)
	if err1 !=nil {
		fmt.Println("os.Create err1 = ", err1)
		return
	}
	defer f.Close()

	f.WriteString(result) //写内容
	page <- i	//堵塞模式
}

func DoWork(start, end int) {
	fmt.Printf("正在爬取%d到%d的页面\n", start, end)

	page := make(chan int)
	//明确目标(要知道你准备在那个范围或者网站去搜索)
	//https://tieba.baidu.com/f?kw=%E7%BB%9D%E5%9C%B0%E6%B1%82%E7%94%9F&ie=utf-8&pn=0
	for i:=start; i<=end; i++ {
		go SpiderPape(i, page)	//新建协程并发爬取
	}

	for i:=start; i<=end; i++ {
		fmt.Printf("第%d页面爬取完成\n", <-page) //爬取完了结束, 主程序收到后结束程序
	}
}

func main()  {

	var start, end int
	fmt.Println("请输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入结束页(>=1):")
	fmt.Scan(&end)

	DoWork(start, end)
}