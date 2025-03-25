# go中的channel

什么是 Channel？
channel 是 Go 语言中用于 goroutine 之间通信 的工具。
你可以把它想象成一个 管道，一边放数据（发送），另一边取数据（接收）。
它是 Go 并发编程的核心，解决了多线程通信时常见的复杂同步问题。

**不要通过共享内存来通信，而要通过通信来共享内存**

来看一段程序
```go main.go
package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 100
	values := <-ch
	close(ch)
	fmt.Println(values)
}
```
直接死锁了，死锁原因是什么呢


因为创建了一个无缓冲区的channel，所以在向管道添加数据的时候要等待其他的`goroutines`读取该值，锁才会被释放，所以一直没有`goroutines`接收值，就会导致死锁。
然后有个疑问就是下面我明明定义了value去接收值，但是就是发生了死锁，原因是我当前`goroutines`已经被阻塞了，走不到value被接收的那一步

好比接力棒，无缓冲的话必须得有人接棒才能跑，不能自己自己跑完还是自己接棒，有缓冲就可以，一个人自己来，而且缓冲区代表能允许多少个人能独立来回跑的过程（典型牛马）


对应的解决方法是：设置`channel`是有缓冲的，或者说定义一个goroutines去接收传入管道的值
