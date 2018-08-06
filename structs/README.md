To update information in a struct we can either update it directly by direct reference

or if we want to use receivers we have to use a pointer because the receivers will create

a new copy of the variable.

### Pointers

&variable ==> give me the memory address of the value the variable is pointing at

*variable ==> give me the value this memory address is pointing at

Here is the rule 
```text
turn address to value using *address and value to address using &value

*ATV&VTA

```


```text
package main
func (pointerToPerson *person) updateFirstName(newFirstname string) {
	(*pointerToPerson).firstName = newFirstname;
}

```

In the above code inside receiver the *person is the type description, it means we are working with a pointer to a person.

The *pointerToPerson means we wnt to manipulate the value the pointer is referencing.


###Note

1) Whenever you pass an integer, float, string, or struct into a function, go creates copy of each argument in the function

2) Go has two types of data types
```text

value type | reference type
---------------------------
int        | slices
float      | maps
string     | pointers
bool       | functions
structs    | channels

So for changing value types we need to use pointers but for reference types we don't need to.
```

So the difference between two types is reference types all actually kind of references in memory to the original values.
FOr eg slice is a pointer to an array so within a function when go creates a copy of the slice reference 
it is still referring to the same array
