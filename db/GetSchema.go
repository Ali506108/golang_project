package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

func main() {

	arguments := os.Args
	if len(arguments) != 6 {
		fmt.Println("please provide hostname and password db  ")
		return
	}

	host := arguments[1]
	port := arguments[2]
	user := arguments[3]
	pass := arguments[4]
	dbName := arguments[5]

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error in opening and connection to db : %v \n", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("in a sadnes we dont connect to db %v \n", err)
	}

	fmt.Println("Successfully connections to database and starting testing db")
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	db.SetConnMaxLifetime(time.Minute * 50)

	workers := 50
	rps := 10000

	var wg sync.WaitGroup
	start_Time := time.Now()

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerId int) {
			defer wg.Done()

			for j := 0; j < rps; j++ {
				_, err := db.Exec("SELECT 1")

				if err != nil {
					log.Printf("Worker %v watched the mistake : %v\n", workerId, err)
				}
			}
		}(i)
	}

	wg.Wait()

	duration := time.Since(start_Time)
	total_req := workers * rps

	fmt.Println("=====================================")
	fmt.Printf("Тест завершен!\n")
	fmt.Printf("Всего выполнено запросов: %d\n", total_req)
	fmt.Printf("Затраченное время: %v\n", duration)
	fmt.Printf("RPS (запросов в секунду): %.2f\n", float64(total_req)/duration.Seconds())
	fmt.Println("=====================================")
}

/*
PS C:\Users\Legion5\GolandProjects\awesomeProject3> go run .\db\GetSchema.go localhost 5433 root root db
Successfully conection to database and starting testing db
=====================================
Тест завершен!
Всего выполнено запросов: 500000
Затраченное время: 1m3.6244245s
RPS (запросов в секунду): 7858.62
=====================================
*/
