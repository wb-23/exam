package main
import "fmt"

var channel = make(chan int,1)
var over=make(chan bool,1)
var c int
func p(){
	c=<-channel
	fmt.Printf("hello world %d\n",c)
	c++
	channel<- c
	if c == 10 {
		over<-true
	}
}
func main(){
	channel<-1
	go p()
	go p()
	go p()
	go p()
	go p()
	go p()
	go p()
	go p()
	go p()
	go p()
	<-over
	<-channel
}

