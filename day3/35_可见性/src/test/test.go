package test

import "fmt"

// 如果想使用倍的函数,结构体类型,结构体成员,函数名,类型名,结构体成员变量名,手写字母为大写时 , 表示可见
// 如果首写字母时小写, 只能再同一个包里使用(即同级目录下引用),表示不可见

func myFunc()  {
	 fmt.Println("this is myFunc") //当为小写字母开头时只能同目录下可调用该函数
}

func MyFunc()  {
	fmt.Println("this is MyFunc") //当为大写字母开头时该工程可调用
}

type stu struct {	//当为小写字母开头时只能同目录下可调用该函数
	id int			//当为小写字母开头时只能同目录下可调用该函数
}

type Stu struct { //当为大写字母开头时该工程可调用
	Id int			//当为大写字母开头时该工程可调用 , 这个里面的Id也是一样的
}