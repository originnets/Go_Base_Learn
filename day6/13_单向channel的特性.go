package main


func main() {
	// 创建一个channel, 默认是双向的
	ch := make(chan int)

	// 双向channel能隐式转换伟单向channel
	var WriteCh chan <- int = ch	//只能写不能读

	WriteCh <- 666 //写
	//<- WriteCh //err invalid operation: <-WriteCh (receive from send-only type chan<- int)

	var ReadCh <- chan int = ch // 只能读不能写

	<-ReadCh // 读

	//ReadCh <- 666 //err invalid operation: ReadCh <- 666 (send to receive-only type <-chan int)

	// 单向是无法转换为双向
	//var ch2 chan int = WriteCh //err cannot use WriteCh (type chan<- int) as type chan int in assignment
}
