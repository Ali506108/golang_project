package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

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

// Напишите функцию, которая запускает несколько горутин, ожидает их выполнения,
// выводит на экран текст об окончании и завершается. При этом, если функция
// ожидает завершения больше двух секунд, то она должна вернуть ошибку.
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Data sent to Product-service")
}

func runGoroutine(n int) error {
	// create waitGroup for telling our main routine is task is done
	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i < n; i++ {
		go worker(i, &wg)
	}

	// We fix here main trouble wite time in a first time i was think about use time.NewTimer() but reviewing with perplexity help me to understund me main trouble
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// wait to	completion our threads	or timeOut
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("All tasks is done")
		return nil
	case <-ctx.Done():
		return fmt.Errorf("Timeout waiting for goroutine took more than 2 seconds")

	}
}

func main() {
	fmt.Println("Start our production")

	if err := runGoroutine(5); err != nil {
		fmt.Println("Error: ", err)
	}
}
