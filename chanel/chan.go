package main

import (
	"fmt"
	"time"
)

func readChan(chanel chan int) {
	value := <-chanel
	fmt.Println("Chanel value is : ", value)
}

func writeChanel(cch chan<- int) {
	for i := 0; i < 10; i++ {
		value := i*i ^ 2
		cch <- value
	}
}

//Напишите функцию foo(), которая обрабатывает панику и возвращает ошибку. В
//теле функции также нужно вызвать панику. Вызовите функцию foo(), проверьте на
//ошибки и выведите на экран текст ошибки.

func foo() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Recovered from panic : %v ", r)
		}
	}()

	panic("Foo trouble")

	return err
}

//Напишите функцию, которая запускает несколько горутин, ожидает их выполнения,
//выводит на экран текст об окончании и завершается. При этом, если функция
//ожидает завершения больше двух секунд, то она должна вернуть ошибку.
func goRoutineFunction() {
	
}

func readChanel(cch, quite <-chan int) {
	for {
		select {
		case x := <-cch:
			fmt.Printf("The value is : %v\n", x)
		case <-quite:
			return
		default:
			fmt.Println("Default value")
			return
		}
	}
}

func main() {
	fmt.Println("Start main channel")

	var ch chan int
	ch = make(chan int, 2)

	ch <- 43
	go readChan(ch)

	ch <- 41

	fmt.Println("End main function")

	quite := make(chan int)
	chanel_main := make(chan int)

	go readChanel(chanel_main, quite)

	go writeChanel(chanel_main)
	time.Sleep(1 * time.Second)

}
