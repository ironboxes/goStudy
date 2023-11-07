package main

import (
	"fmt"
	"time"
)

type Worker struct {
	stream  <-chan int    //处理
	timeout time.Duration //超时
	done    chan struct{} //结束信号
}

func NewWorker(stream <-chan int, timeout int) *Worker{
	return &Worker{
		stream: stream,
		timeout: time.Duration(timeout)*time.Second,
		done: make(chan struct{}),
	}
}

func (w *Worker) afterTimeStop() {
	go func ()  {
		time.Sleep(w.timeout)
		w.done <- struct{}{}
	}()
}
func (w *Worker) Start() {
	w.afterTimeStop()
	for {
		select {
		case data, ok := <-w.stream:
			if !ok {
				return
			}
			fmt.Println(data)
		case <-w.done:
			close(w.done)
			return
		}
	}
}

func main() {
	//test()
	test1()
}

func test() {
	stream := make(chan int)
	defer close(stream)

	w := NewWorker(stream, 3)
	w.Start()
}

func test1() {
	    done := make(chan struct{})
timer := time.NewTimer(1 * time.Second)
go func () {
for {
select {
case <-timer.C:
fmt.Println(time.Now())
timer.Reset(1 * time.Second)
case <-done:
return
}
}
}()
<-time.After(5*time.Second + time.Millisecond*100)
done <- struct{}{}
}