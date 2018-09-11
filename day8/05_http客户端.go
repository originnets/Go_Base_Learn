package main

import (
	"fmt"
	"net/http"
)

func main()  {
	//resp, err := http.Get("http://www.baidu.com")
	resp, err := http.Get("http://127.0.0.1:8000")
	if err != nil {
		fmt.Println("http.Get err = ", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("resp.Status = ", resp.Status)
	fmt.Println("resp.StatusCode = ",resp.StatusCode)
	fmt.Println("resp.Header = ",resp.Header)
	//fmt.Println("resp.Body = ",resp.Body)

	buf := make([]byte, 4*1024)
	var tmp string
	for {
		n, err := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("resp.Body.Read err = ", err)
			break
		}
		tmp += string(buf[:n])
	}
	fmt.Println("body = ", tmp)
}