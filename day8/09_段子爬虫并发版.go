package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func HttpGet(url string)(result string, err error){
	resp, err1 := http.Get(url) //发送请求
	if err1 != nil{
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网页内容

	buf := make([]byte, 4*1024)
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			if err == io.EOF{
				//读取结束
			}else {
				fmt.Println("resp.Body.Read err = ", err)
			}
			break
		}
		result += string(buf[:n]) //累加读取的内容
	}
	return
}

func StoreJoyToFile(i int, fileTile, fileContent []string) {
	//新建文件
	f, err := os.Create(strconv.Itoa(i) + ".txt")
	if err != nil {
		fmt.Println(" os.Create err = ", err)
		return
	}

	defer f.Close()

	//写内容

	n := len(fileTile)

	for i:=0;i<n ;i++  {
		//写标题
		f.WriteString(fileTile[i] + "\n")
		//写内容
		f.WriteString(fileContent[i]+ "\n")

		f.WriteString("\n================================================分割线=========================================\n")
	}
}

//爬去title,content
func SpiderOnerJoy(url string) (title, content string, err error){
	//开始爬取页面内容
	result, err1 := HttpGet(url)
	if err1 != nil {
		//fmt.Println("HttpGet err", err)
		err = err1
		return
	}
	//去关键信息
	//取段子标题  <h1> 标题 </h1>

	re1 := regexp.MustCompile(`<h1>(?s:(.*?))</h1>`)
	if re1 == nil{
		//fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	Tmptitle := re1.FindAllStringSubmatch(result, 1) //最后一个参数为1, 只过滤一次
	for _, data := range Tmptitle {
		//fmt.Println("title = ", data[1])
		buf := strings.Fields(data[1]) //去除两边的空格,返回的时切片
		for _, data := range buf {
			title += data	//将切片转换为字符串
			title = strings.Replace(title, "<br/>", "",-1) //去除<br/>标签
			title = strings.Replace(title, "&nbsp;", " ",-1) //把&nbsp;换成空格
		}
		break
	}
	//取段子内容  <div class="content-txt pt10"> 段子内容 <a id="prev" href="
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil{
		//fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}
	Tmpcontent := re2.FindAllStringSubmatch(result, 1) //最后一个参数为1, 只过滤一次
	for _, data := range Tmpcontent {
		//fmt.Println("content = ", data[1])
		buf := strings.Fields(data[1])	//去除两边的空格,返回的时切片
		for _, data := range buf {
			content += data	//将切片转换为字符串
			content = strings.Replace(content, "<br/>", "",-1)
			content = strings.Replace(content, "&nbsp;", " ",-1)
		}
		break
	}
	return
}


//爬取主页面url
func SpiderPape(i int, page chan int) {
	//明确爬取的url
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) +".html"
	fmt.Printf("正在爬取第%d的网页:%s\n", i, url)

	//开始爬取页面内容
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err", err)
		return
	}
	//fmt.Println("r = ", result)
	//取段子url内容 <h1 class="dp-b"><a href="段子url " target="_blank">
	//解析表达式
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))" target="_blank"`)
	if re == nil{
		fmt.Println("regexp.MustCompile err")
		return
	}
	//取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1) //过滤所有URl

	fileTile := make([]string, 0)
	fileContent := make([]string, 0)

	//fmt.Println("joyUrls = ", joyUrls)
	//取网址
	//第一个返回下标,第二个返回内容
	for _,data := range joyUrls{
		//fmt.Println("url =", data[1])
		//开始爬去每一个段子
		title, content, err := SpiderOnerJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOnerJoy err = ", err)
			continue
		}
		//fmt.Println("title = ", title)
		//fmt.Println("content = ", content)
		fileTile = append(fileTile, title) 	//追加标题
		fileContent = append(fileContent, content)	//追加内容
	}
	//fmt.Println("fileTile = ", fileTile)
	//fmt.Println("fileContent = ", fileContent)

	//把内容写入文件
	StoreJoyToFile(i,fileTile, fileContent)
	page <- i	//写内容, 写num
}

func DoWork(start, end int) {
	fmt.Printf("准备爬取%d到%d的页面\n", start, end)

	page := make(chan int)
	for i:= start; i<=end; i++ {
		//定义一个函数,爬去主页面
		go SpiderPape(i, page)
	}
	for i:= start; i<=end; i++ {
		fmt.Printf("第%d页面爬取完成\n", <-page)
	}
}

func main()  {
	var start, end int
	fmt.Println("请输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Println("请输入结束页(>=1):")
	fmt.Scan(&end)

	DoWork(start, end) //工作函数
}
