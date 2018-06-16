package main

import "fmt"

func abstractListner(fxChan chan func()string){
	fxChan<- func() string {
		return "sent"
	}
}

func main() {
	fxChan:=make(chan func() string)
	defer close(fxChan)
	go abstractListner(fxChan)
	select{
	  case rfx:=<-fxChan:
	  	msg:=rfx()
	  	fmt.Println(msg)
	  	fmt.Println("Received")
	}
}
