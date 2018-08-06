package main

import (
	"net/http"
	"fmt"
	"os"
)

type logWriter struct {}
func readFile(resp http.Response){
	bs:=make([]byte,99999);
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

func main() {
	resp,err:=http.Get("https://blog.deepanshululla.com")
	if (err!=nil){
		fmt.Println("Error",err)
		os.Exit(1)
	}
	fmt.Println(*resp)
	//readFile(*resp);

	// this works too
	//io.Copy(os.Stdout,resp.Body)

	// custom implementation of writers interface
	//lw:=logWriter{}
	//io.Copy(lw,resp.Body)


}

func (logWriter) Write(bs[] byte) (int,error){
	fmt.Println(string(bs))

	return len(bs),nil
}