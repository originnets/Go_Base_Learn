package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	list := os.Args //获取命令行参数

	if len(list) != 3 {
		fmt.Println(" usage: xxx SrcFile DstFile")
		return
	}
	SrcFileName := list[1]
	DstFileName := list[2]

	if SrcFileName == DstFileName {
		fmt.Println("源文件和目的文件名字不能相同")
		return
	}

	// 只读方式打开源文件
	SF , err1 := os.Open(SrcFileName)
	if err1 != nil {
		fmt.Println("err1 = ", err1)
		return
	}

	//新建目的文件
	DF , err2 := os.Create(DstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//操作完毕, 需要关闭文件
	defer DF.Close()
	defer SF.Close()

	//核心处理, 从源文件读取内容,往目的文件写, 读多少写多少
	buf := make([]byte, 4*1024)	//4K大小历史缓冲区
	for {
		r, rerr := SF.Read(buf) //从源文件读取内容,写到buf中 , 再返回一个buf使用index=r
		if rerr != nil {
			if rerr == io.EOF{ //文件读取完毕
				break
			}
			fmt.Println("rerr = ", rerr)
		}
		//往目的文件写
		DF.Write(buf[:r])
	}
}
