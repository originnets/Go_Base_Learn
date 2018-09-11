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

func main()  {
	//定义一个接口类型的变量

	var i Personer
	s := &Student{"mirk", 111}
	i = s
	i.SayHi() // 继承的方法
	i.Sing("歌歌歌")

}