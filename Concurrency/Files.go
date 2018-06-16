package main

import (
	"sync"
	"io/ioutil"
	"strconv"
	"fmt"
)

var writer chan bool
var rwLock sync.RWMutex

func writeFile(i int){
	rwLock.RLock()
	ioutil.WriteFile("test.txt",
		             []byte(strconv.FormatInt(int64(i),
			         10)),0x777)
	rwLock.RUnlock()
	writer<-true
}


func main() {
	writer=make(chan bool)
	for i:=0;i<10;i++{
		go writeFile(i)
		fmt.Println(i)
	}
	<-writer
	fmt.Println("Done")
}
