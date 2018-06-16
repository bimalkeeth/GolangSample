package main

import (
	"strconv"
	"fmt"
)

type Messanger interface {

	Relay()string
}
type Message struct{
	status string
}
func(m Message) Relay()string{
	return m.status
}

func alertMessage(v chan Messanger,i int){
	m:=new(Message)
	m.status="Done with " +strconv.FormatInt(int64(i),10)
	v<-m
}

func main() {
	msg:=make(chan Messanger)
	for i:=0;i<10;i++{
		go alertMessage(msg,i)
	}
	select{
	  case message:=<-msg:
	  	fmt.Println(message.Relay())
	}
	<-msg
}
