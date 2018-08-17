package main

import (
	"fmt"
)



func opsWorksChannelInstanceName(fileName string, c chan string){

	opsWorksStructGrp:=InitServiceGroup()
	fmt.Println("Aws Call for ec2")
	l:=opsWorksStructGrp.GetInstanceIdNameMap()
	fmt.Println(" From writing: ",len(l))
	ConvertToJsonFile(CreateStructMapForMap(l),fileName)
	fmt.Println("Successfuly wrote to file")
	c<-fileName
}


func opsWorksChannelElbStackName(fileName string, c chan string){
	fmt.Println("AWS call for elb")
	opsWorksStructGrp:=InitServiceGroup()
	l:=opsWorksStructGrp.GetELbStackNameMap()
	fmt.Println(" From writing: ",len(l))
	ConvertToJsonFile(CreateStructMapForMap(l),fileName)
	fmt.Println("Successfuly wrote to file")
	c<-fileName
}


func syncher(){
	fileName1:="instanceIdNameMap.json"
	fileName2:="elbClusterNameMap.json"
	chan1:=make(chan string)
	chan2:=make(chan string)

	go opsWorksChannelInstanceName(fileName1,chan1)
	go opsWorksChannelElbStackName(fileName2,chan2)
	for{
		select {
		case _=<-chan1:
			go opsWorksChannelInstanceName(fileName1,chan1)
		case _=<-chan2:
			go opsWorksChannelElbStackName(fileName2,chan2)
		default:
			dict:=ReadJsonMap(fileName1)
			fmt.Println("Length after reading file1",len(dict))
			dict2:=ReadJsonMap(fileName2)
			fmt.Println("Length after reading file2",len(dict2))

		}
	}

}

func main(){
	syncher()
}

