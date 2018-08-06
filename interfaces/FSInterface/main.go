package main

import (
	"fmt"
	"os"
	"io"
)
type logWriter struct {}
func (logWriter) Write(bs[] byte) (int,error){
	fmt.Println(string(bs))

	return len(bs),nil
}


func main() {
	fName:=os.Args[1];
	f,error:=os.Open(fName);
	if (error!=nil){
		fmt.Println(error)
		os.Exit(1)
	}
	//bs:=make([]byte,99999);
	io.Copy(os.Stdout,f)
}
