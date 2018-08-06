package main

import (
	"net/http"
	"fmt"
	"time"
)

func checkLink(link string)  {
	_,err:=http.Get(link)
	if err!=nil{
		fmt.Println(link, "is down")
		return
	}
	fmt.Println(link,"is up")

}

func serialLinkChecker(links[] string){
	fmt.Println("Checking status of websites serially ")
	for _,link := range links {
		checkLink(link)
	}
}

func checkLinkRoutine(link string, c chan string) {

	_,err:=http.Get(link)
	if err!=nil{
		fmt.Println(link, "is down")
		c<-link
		return
	}
	fmt.Println(link,"is up")
	c<-link

}

func goRoutineLinkChecker(links[] string){
	fmt.Println("Checking status of websites using go routines ")
	//creating a channel
	c:=make(chan string)

	for _,link := range links {
		go checkLinkRoutine(link,c);
	}


	//for i:=0;i<len(links);i++{
	//	fmt.Println(<- c)
	//}
	//for {
	//	go checkLinkRoutine(<- c, c);
	//}
	//for l:= range c{
	//	go checkLinkRoutine(l, c);
	//}
	for l:= range c{
		go func(link string) {
			checkLinkRoutine(link, c);
			time.Sleep(10*time.Second)
		}(l);
	}
}

func main() {
	links:=[]string{"https://www.facebook.com",
				"http://www.deepanshululla.com",
				"https://blog.deepanshululla.com",
				"https://www.google.com",
				//"https://nonexistingwebsite.com",
				}
	//serialLinkChecker(links)
	goRoutineLinkChecker(links)
}
