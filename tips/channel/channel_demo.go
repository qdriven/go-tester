package channel

import (
	"fmt"
	"time"
)

type GenericDataEvent struct {
	Data interface{}
}

/**
ChannelType = ( "chan" | "chan" "<-" | "<-" "chan" ) ElementType .
chan T          // 可以接收和发送类型为 T 的数据
chan<- float64  // 只可以用来发送 float64 类型的数据
<-chan int      // 只可以用来接收 int 类型的数据
如果没有设置容量，或者容量设置为0, 说明Channel没有缓存，只有sender和receiver都准备好了后它们的通讯(communication)才会发生(Blocking)。
如果设置了缓存，就有可能不发生阻塞， 只有buffer满了后 send才会阻塞， 而只有缓存空了后receive才会阻塞。一个nil channel不会通信。

go routine: send/receive
Channel可以作为一个先入先出(FIFO)的队列，接收的数据和发送的数据的顺序是一致的。
v, ok := <-ch
*/
func MakeChannel() chan *GenericDataEvent {
	ch := make(chan *GenericDataEvent, 100)
	return ch
}

func SumByChannelWay(items []*GenericDataEvent, c chan *GenericDataEvent) {
	sum := 0
	for _, item := range items {
		data, ok := item.Data.(int)
		if !ok {
			fmt.Println("convert to int failed")
		}
		sum += data
	}
	c <- &GenericDataEvent{
		Data: sum,
	}
}

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- time.After(time.Second*1):
			fmt.Println("timeout 1")
		case <-quit:
			fmt.Println("quit loop signal")
			return
		}
	}
}
func ExecFib(num int) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < num; i++ {
			fmt.Println("sending data.......")
			time.Sleep(time.Second*2)
			fmt.Println(<-c)
		}
		quit <- 0 //quite function
	}()
	Fibonacci(c, quit)
}

func ExecuteSum() {
	c := MakeChannel()
	s := make([]*GenericDataEvent, 0)
	for i := 0; i < 10; i++ {
		s = append(s, &GenericDataEvent{Data: i})
	}
	go SumByChannelWay(s[:len(s)/2], c)
	go SumByChannelWay(s[len(s)/2:], c)
	x, y := <-c, <-c
	xInt, _ := x.Data.(int)
	yInt, _ := y.Data.(int)
	fmt.Println(x, y, xInt+yInt)
}

func SendToChannel() {
	ch := MakeChannel()
	go func() {
		ch <- &GenericDataEvent{Data: 3}
	}()
	defer close(ch)
}

//**往一个已经被close的channel中继续发送数据会导致run-time panic。
func SendToChannelPanic() {
	ch := MakeChannel()
	close(ch)
	go func() {
		ch <- &GenericDataEvent{Data: 3}
	}()
}

//TIMER/TICKER

func TimerTickerUsage(){
	timer1 :=time.NewTimer(time.Second*2)
	<-timer1.C
	fmt.Println("Timer1 expired")
}

//https://colobu.com/2016/04/14/Golang-Channels/
//http://legendtkl.com/2017/07/30/understanding-golang-channel/