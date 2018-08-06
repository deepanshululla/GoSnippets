package main

import "fmt"

type student struct {
	firstName string
	lastName  string
}
type contactInfo struct {
	email   string
	zipcode string
}

//type person struct {
//	firstName string
//	lastName string
//	contact contactInfo
//}
type person struct {
	firstName string
	lastName  string
	contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

/*
This will not work because go will create a new copy of the variable
and update the copy. it will not update the original


func (p person) updateFirstName(newFirstname string){
	p.firstName=newFirstname;
}

 */

/*
Neither this will work

func updFName(p person,newFirstname string){
	p.firstName=newFirstname
}
 */

/*
To update the original entry we need to use pointers

 */
func (pointerToPerson *person) updateFirstName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname;
}

func basicStructs() {
	deepanshu := student{"deepanshu", "lulla"}
	deepu := student{firstName: "deepu", lastName: "lulla"}
	var alex student
	alex.firstName = "Alex"
	fmt.Printf("%+v\n", deepanshu)
	fmt.Println(deepanshu)
	fmt.Printf("%+v\n", deepu)
	fmt.Println(deepu)
	fmt.Printf("%+v\n", alex)
	fmt.Println(alex)
}

func embedStructsWithinStructs() {
	alex := person{firstName: "Alex",
		lastName: "farrow",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipcode: "90533",
		}};
	fmt.Printf("%+v\n", alex)
	fmt.Println(alex)
	// this works too
	//alexPointer := &alex;
	//alexPointer.updateFirstName("alexander")
	//this will work as well
	alex.updateFirstName("alexander")
	alex.print()


}
/*
The below function works because slice is a reference type

 */

func updateSlice(s []string){
	s[0]="bye";
}

func main() {

	embedStructsWithinStructs()
	newSlice:=[]string{"hi","there","how","are","you"}
	fmt.Println(newSlice)
	updateSlice(newSlice)
	fmt.Println(newSlice)
}
