package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	arg := os.Args

	if len(arg) != 6 {
		fmt.Errorf("Please provide hostname port password username dbName")
	}

	host := arg[1]
	port := arg[2]
	username := arg[3]
	password := arg[4]
	dbName := arg[5]

	ConnectionStr := fmt.Sprintf("please type own host=%s port=%s username=%s password=%s dbName=%s , sslmode=disable ",
		host, port, username, password, dbName,
	)

	db, err := sql.Open("postgres", ConnectionStr)
	if err != nil {
		fmt.Errorf("Error opening and connection to Db : %v \n", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Errorf("We have trouble to connection to db : %v ", err)
	}

	fmt.Println("Successfully connect to Database and starting test ")
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Minute * 5)

	workers := 50
	rps := 10000

	var wg sync.WaitGroup
	start_time := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()

			for j := 0; j < rps; j++ {
				_, err := db.Exec("SELECT 1")

				if err != nil {
					fmt.Errorf("worker %v find the mistake : %v ", j, err)
				}
			}
		}(i)
	}

	wg.Wait()

	duration := time.Since(start_time)
	total_req := workers * rps

	fmt.Println("=====================================")
	fmt.Printf("Тест завершен!\n")
	fmt.Printf("Всего выполнено запросов: %d\n", total_req)
	fmt.Printf("Затраченное время: %v\n", duration)
	fmt.Printf("RPS (запросов в секунду): %.2f\n", float64(total_req)/duration.Seconds())
	fmt.Println("=====================================")

}
