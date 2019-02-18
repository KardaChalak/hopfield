# hopfield


A (unoptimized) go version of the project described on <a href=http://chalak.se/> http://chalak.se/ </a> (Projects -> Associative Memory)
This is a one day hack to start learning golang.

A hopfield network can be used to correct distorted patterns. If we have a network with 1000 neurons it can remember around 138.
The exact amount depends on the composition of the patterns.


The package is easy to use.
It contains 2 methods and a "constructor" function ( + a String() string method).

* func NewNet(nrNeurons int) (Net, error)
* func (net Net) InsertPattern(pattern []int) error
* func (net Net) Recall(pattern []int) ([]int, error)

If the possible errors are disregarded we get this example:


```go
package main

import (
	"fmt"
	"hopfield"
)

func main() {
	pattern1 := []int{1, 1, 1, 1, 1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1}
	pattern2 := []int{1, 1, 1, 1, 1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}
	pattern3 := []int{1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, -1, -1, -1, -1, 1, 1, 1, 1, 1}

	distortedPattern := []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1}
	net, _ := hopfield.NewNet(25)
	net.InsertPattern(pattern1)
	net.InsertPattern(pattern2)
	net.InsertPattern(pattern3)
	correctedPattern, _ := net.Recall(distortedPattern)

	fmt.Println(correctedPattern)

}
}
```

First a hopfield net with 25 neurons is created. Then 3 patterns are insearted. Lastly the ditorted pattern is corrected and prints:

```
[1 -1 -1 -1 -1 1 -1 -1 -1 -1 1 -1 -1 -1 -1 1 -1 -1 -1 -1 1 1 1 1 1]
```
Which is pattern3

