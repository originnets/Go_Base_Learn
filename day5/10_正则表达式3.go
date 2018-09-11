package main

import (
	"fmt"
	"regexp"
)
// ``反引号表示原生字符串
func main() {

	buf := `
	<!DOCTYPE html>
<html lang="zh-CN">
<head>
	<title>Go语言标准库文档中文版 | Go语言中文网 | Golang中文社区 | Golang中国</title>
	<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1.0, user-scalable=no">
	<meta http-equiv="X-UA-Compatible" content="IE=edge, chrome=1">
	<meta charset="utf-8">
	<link rel="shortcut icon" href="/static/img/go.ico">
	<link rel="apple-touch-icon" type="image/png" href="/static/img/logo2.png">
	<meta name="author" content="polaris <polaris@studygolang.com>">
	<meta name="keywords" content="中文, 文档, 标准库, Go语言,Golang,Go社区,Go中文社区,Golang中文社区,Go语言社区,Go语言学习,学习Go语言,Go语言学习园地,Golang 中国,Golang中国,Golang China, Go语言论坛, Go语言中文网">
	<meta name="description" content="Go语言文档中文版，Go语言中文网，中国 Golang 社区，Go语言学习园地，致力于构建完善的 Golang 中文社区，Go语言爱好者的学习家园。分享 Go 语言知识，交流使用经验">
</head>
	<div>哈哈</div>
	<div>
	sadfasasdsa
	sdsad
	</div>
	<div>呵呵</div>
	<div>阿迪斯撒旦</div>
	<div>阿迪斯</div>

<frameset cols="15,85">
	<frame src="/static/pkgdoc/i.html">
	<frame name="main" src="/static/pkgdoc/main.html" tppabs="main.html" >
	<noframes>
	</noframes>
	</frameset>
</html>
`
	//定义一个正则表达式
	//reg := regexp.MustCompile(`<div>(.*)</div>`)
	reg := regexp.MustCompile(`<div>(?s:(.*?))</div>`)

	if reg == nil {
		fmt.Println("error")
		return
	}

	//提取关键信息

	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)	//数组的方式打印
	//fmt.Println("result = ", result)

	// 过滤器<></>
	for _, text := range result {
		//fmt.Println("text[0] = ", text[0])	//带<></>
		fmt.Println("text[1] = ", text[1])	//不带<></>
	}

}
