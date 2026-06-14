package main

import (
	"fmt"
	"time"
)

func getData() {
	panic("Panic in get Data")
}

func main() {
	defer func() {
		recEver := recover()
		fmt.Println("restart the system")
		fmt.Println(recEver)
	}()
	fmt.Println("Data is ", time.Now())
	getData()
	fmt.Println("Data")
}
