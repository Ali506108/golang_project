package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	runtime.GOMAXPROCS(1)

	// single responsibility
	//Dry ,concise, low-key humor
	// goroot , gopath,goproxy

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i)
	}

	wg.Wait()
}
