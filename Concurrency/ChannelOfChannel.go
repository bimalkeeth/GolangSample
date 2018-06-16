package main

import "fmt"

type Item struct{
	Url string
	Data []byte
}
type Feed struct{
	Url string
	Name string
	Items []Item
}

type master chan Item
var Feeds[]Feed

func process(feedChannel *chan master,done *chan bool ){
	for _,i:=range Feeds{
		fmt.Println("feed",i)
		item:=Item{}
		item.Url=i.Url
		itemChannel:=make(chan Item)
		*feedChannel<-itemChannel
		itemChannel<-item
	}
    *done<-true
}

func processItem(url string){
	fmt.Println("Got url",url)
}
func main() {
	done:=make(chan bool)
	Feeds=[]Feed{{Name: "Newyork Times", Url: "http://rss.nytimes.com/services/xml/rss/nyt/HomePage.xml"},
		{Name: "Wall	Street	Journal", Url:
		"http://feeds.wsjonline.com/wsj/xml/rss/3_7011.xml"}}
	   feedChannel:=make(chan master)
	   go func(done chan bool,feedChannel chan master){
	   	 for{
	   	 	select{
	   	 	   case fc:=<-feedChannel:
	   	 	   	select{
	   	 	   	   case item:=<-fc:
					   processItem(item.Url)
				}
			default:

			}
		 }
	   }(done,feedChannel)
	   go process(&feedChannel,&done)
	   <-done
	   fmt.Println("Done")
}
