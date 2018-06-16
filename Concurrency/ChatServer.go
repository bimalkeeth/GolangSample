package main

import (
	"net"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)
var connectionCount=0
var messagePool chan(string)
const(INPUT_BUFFER_LENGTH=140)
var	Usersx	[]*UserX

type UserX struct{
	Name string
	ID int
	Initiated bool
	UChannel chan []byte
	Connection *net.Conn
}
func(u UserX) Listen(){
	fmt.Println("Listning for",u.Name)
	for{
		select{
		case msg:=<-u.UChannel:
			fmt.Println("Sending new message to ",u.Name)
			fmt.Fprintln(*u.Connection,string(msg))
		}
	}
}
type ConnectionManager struct{
	name string
	initiated bool
}
func Initiated() *ConnectionManager{
	cM:=&ConnectionManager{name:"Chat Server 1,o",initiated:false}
	return cM
}
func evalMessageReceipent(msg []byte,uName string) bool{
	eval:=true
	expression:="@"
	re,err:=regexp.MatchString(expression,string(msg))
	if err!=nil{
		fmt.Println("Error:",err)
	}
	if re==true{
		eval=false
		pmExpression:="@"+uName
		pmRe,pmErr:=regexp.MatchString(pmExpression,string(msg))
		if pmErr!=nil{
			fmt.Println("Regex Error ",err)
		}
		if pmRe==true{
			eval=true
		}
	}
	return eval
}
func(cm *ConnectionManager) Listen(listener net.Listener){
	fmt.Println(cm.name,"Started")
	for{
		conn,err:=listener.Accept()
		if err!=nil{
			fmt.Println("Connection in error",err)
		}
		connectionCount++
		fmt.Println(conn.RemoteAddr(),"connected")
		user:=UserX{Name:"anonymous",ID:0,Initiated:false}
		Usersx=append(Usersx,&user)
		for _,u :=range Usersx{
			fmt.Println("User On lien",u.Name)
		}
		fmt.Println(connectionCount," connections active")
		go cm.messageReady(conn,&user)
	}
}
func (ma *ConnectionManager) messageReady(conn net.Conn, x *UserX) {
	uChan:=make(chan []byte)
	for{
		buf:=make([]byte,INPUT_BUFFER_LENGTH)
		n,err:=conn.Read(buf)
		if err!=nil{
			conn.Close()
			conn=nil
		}
		fmt.Println(n,"charachter message from user",x.Name)
		if x.Initiated==false{
			fmt.Println("New user is",string(buf))
			x.Initiated=true
			x.UChannel=uChan
			x.Name=string(buf[:n])
			x.Connection=&conn
			go x.Listen()
			minusYouCount:=strconv.FormatInt(int64(connectionCount-1),10)
			conn.Write([]byte("Welcome to the chat,"+x.Name+",there are "+minusYouCount+ "other users"))
		}else{
			sendMessage:=[]byte(x.Name +": " + strings.TrimRight(string(buf),"\t\r\n"))
			for _,u :=range Usersx{
				if evalMessageReceipent(sendMessage,u.Name)==true{
					u.UChannel<-sendMessage
				}
			}
		}
	}
}

func main() {
	connectionCount=0
	serverClosed:=make(chan bool)
	listner,err:=net.Listen("tcp",":9000")
	if err!=nil{
		fmt.Println("Could not start server",err)
	}
	connManage:=Initiated()
	go connManage.Listen(listner)
	<-serverClosed
}
