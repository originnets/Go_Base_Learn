package main

import "fmt"

type Humaner interface {//子集
	SayHi()
}

type Personer interface { //超集
	Humaner // 你名字段, 继承了SayHi()
	Sing(lrc string)
}

type Student struct {
	name string
	id int
}


// Student 实现了SayHi()

func (tmp * Student) SayHi() {
	fmt.Printf("Student[%s, %d] SayHi\n", tmp.name, tmp.id)
}

func (tmp *Student) Sing(lrc string){
	fmt.Println("Student在唱着:", lrc)
}



func main() {
	//超级可以转化为子集,反过来不可以转换
	var ipro Personer	//超集
	ipro =  &Student{"make", 333}
	var i Humaner //子集

	//ipro = i //error
	i = ipro //可以转换,超级可以转换为子集
	i.SayHi()
}