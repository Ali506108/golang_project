package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
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

func work_with_concurency(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("[DOWN] %s : %v \n", url, err)
		return
	}
	defer resp.Body.Close()

	fmt.Printf("[%d] %s \n", url, resp.StatusCode)
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

func checkUrlForValid(url string) {
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("[DOWN] %s : %v\n", url, err)
	}
	defer resp.Body.Close()
	fmt.Printf("[%d] %s \n", resp.StatusCode, url)
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "Prefix_" + origin
	length = len(res)

	return res, length
}

func factorialN(n int) int {
	if n < 0 {
		return 1
	}
	return factorialN(n-1) * n
}

func factorialN_for_commit(n int) int {
	if n < 0 {
		return 1
	}

	return factorialN_for_commit(n-1) * n
}

type ar2x2 [2][2]int64

func addNumber(a, b ar2x2) ar2x2 {
	num := ar2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			num[i][j] = a[i][j] + b[i][j]
		}
	}
	return num
}

func (a *ar2x2) addNumber(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			a[i][j] = a[i][j] + b[i][j]
		}
	}
}

func main() {

	//if len(os.Args) != 9 {
	//	fmt.Println("Needed 8 integer")
	//	return
	//}
	//
	//k := [8]int64{}
	//
	//for index, i := range os.Args[1:] {
	//	v, err := strconv.Atoi(i)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	k[index] = int64(v)
	//}
	//a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	//b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}
	//
	//fmt.Println("Traditional a+b ", addNumber(a, b))
	//a.addNumber(b)
	//fmt.Println("a+b", a)

	//if len(os.Args) != 9 {
	//	fmt.Errorf("needed 8 integer")
	//	return
	//}
	//
	//fmt.Println("Traditional a+b ", addNumber(a, b))
	//a.addNumber(b)
	//fmt.Println("a+b", a)

	urls := []string{
		"https://golang.com",
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

	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	for i := 1; i <= 5; i++ {
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

	for key, value := range citys {
		fmt.Printf("Key : %d , value : %d \n", key, value)
	}

	args := map[string]string{
		"Apple":    "software/hardware compony",
		"Google":   "software compony",
		"openai":   "software compony",
		"Barclysr": "Bank compony",
	}

	for _, val := range args {
		fmt.Printf("value %d ", val)
	}

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

	s_data := strings.Map(func(r rune) rune { return r + 1 }, "SDWS")

	fmt.Println(s_data)

	data := adder()
	fmt.Println(data(5))
}

func someFunction() (int, int) {
	return 1, 4
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
