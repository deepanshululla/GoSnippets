#### Wait Groups
The program uses the WaitGroup type of sync package, which is used to wait for the program to finish
all goroutines launched from the main function. Otherwise the goroutines would be launched from main
function and then terminate the program before completing the execution of goroutines. 

The Wait method of the WaitGroup type waits for the program to finish all goroutines. The WaitGroup type uses a counter
that specifies the number of goroutines, and Wait blocks the execution of the program until the WaitGroup
counter is zero.

```text
var wg sync.WaitGroup
wg.Add(2)

```
The Add method is used to add a counter to the WaitGroup so that a call to the Wait method blocks
execution until the WaitGroup counter is zero.
 
 Here a counter of two is added into the WaitGroup , one for
each goroutine.