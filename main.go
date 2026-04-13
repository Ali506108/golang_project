package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

type Movie struct {
	title       string
	description string
	movieRating int
	actors      []Actor
	Studio      string
	director    []string
}

type Actor struct {
	name     string
	sureName string
}

func createMovie() Movie {
	httydMovie := Movie{
		title:       "How to train you're dragon",
		description: "This is the description of the movie",
		movieRating: 10,
		actors: []Actor{
			{
				name:     "John Doe",
				sureName: "John Doe",
			},
			{
				name:     "Jane Doe",
				sureName: "Jane Doe",
			},
		},
		Studio: "DreamWorks",
		director: []string{
			"Steven Spielberg",
		},
	}

	return httydMovie

}

// yeah, you right . I agree with you the true  engineer write own system
// yeah you right ! i agree with you . The true engineer write

// 48 * 24 = 48 * 4 = 192 + 48 * 2 = 96 + 192 = 288

type statusReport struct {
	url    string
	status string
	err    error
}

func checkingUrl(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("[DOWN] %s : %v \n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[%d] %s\n", resp.StatusCode, url)
}

func checkUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("[DOWN] %s : %v\n", url, err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	fmt.Printf("[%d] %s\n", resp.StatusCode, url)
}

func main() {

	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://uber.com",
		"https://this-is-a-fake-url.com",
	}
	start := time.Now()

	for _, url := range urls {
		checkUrl(url)
	}
	fmt.Println("Took ", time.Since(start))
	fmt.Printf("Took %v\n", time.Since(start))

	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>

	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
		//TIP <p>To start your debugging session, right-click your code in the editor and select the Debug option.</p> <p>We have set one <icon src="AllIcons.Debugger.Db_set_breakpoint"/> breakpoint
		// for you, but you can always add more by pressing <shortcut actionId="ToggleLineBreakpoint"/>.</p>
		fmt.Println("i =", 100/i)
	}

	name := "Alex"

	fmt.Println("my name is ", name)
	fmt.Println("Hi that was with logging ", name)

	var citys [8]string
	citys[0] = "New-York"
	citys[1] = "London"
	citys[2] = "Berlin"
	citys[3] = "Los Angeles"
	citys[4] = "Seattle"
	citys[5] = "Austin"
	citys[6] = "Munich"
	citys[7] = "Roma"
	fmt.Println(citys)

	numbers := [10]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(numbers)

	for index, value := range numbers {
		fmt.Printf("index is %d and value is %d \n", index, value)
	}

	for _, v := range numbers {
		fmt.Printf(" Value is %d \n", v)
	}

	valueOne, valueTwo := someFunction()
	fmt.Println("ValueOne is ", valueOne, " value two is ", valueTwo)

	fmt.Println("\nMovie: ", createMovie())

	for i := range 10 {
		fmt.Println("i =", i)
	}

}

func someFunction() (int, int) {
	return 1, 4
}
