Challenge 1
```
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	c <- 1
	fmt.Println(<-c)
}

// This results in a deadlock.
// Can you determine why?
// And what would you do to fix it?

```
Challenge 2
```text
package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	fmt.Println(<-c)
}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?

```