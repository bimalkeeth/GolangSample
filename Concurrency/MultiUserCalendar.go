package main

import (
	"html/template"
	"sync"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type User struct{
	Name string
	Times map[int]bool
	DateHtml template.HTML
}

type Page struct{
	Title string
	Body template.HTML
	Users map[string] User
}

var userInt map[string]bool
var userIndex int
var validTimes []int
var mutex sync.Mutex
var Users map[string]User
var templates=template.Must(template.New("template").ParseFiles("view_users.html","register.html"))

func register(w http.ResponseWriter, r *http.Request){
	fmt.Println("Request to /register")
	params:=mux.Vars(r)
	name:=params["name"]
	if _,ok:=Users[name];ok{

		t,_:=template.ParseFiles("generic.txt")
		page	:=	&Page{	Title:	"User	already	exists",	Body:
		template.HTML("User	"	+	name	+	"	already	exists")}
		t.Execute(w,	page)
	}else {
		newUser	:=	User	{	Name:	name	}
		initUser(&newUser)
		Users[name]	=	newUser
		t,_	:=	template.ParseFiles("generic.txt")
		page	:=	&Page{	Title:	"User	created!",	Body:
		template.HTML("You	have	created	user	"+name)}
		t.Execute(w,	page)
	}
}

func	dismissData(st1	int,	st2	bool)	{
	//	Does	nothing	in	particular	for	now	other	than	avoid	Go	compiler errors
}





func initUser(user *User) {

}
func main() {


}
