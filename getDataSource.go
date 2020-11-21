package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "postgres"
	dbname   = "Reconciliation"
  )

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	  "password=%s dbname=%s sslmode=disable",
	  host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
	  panic(err)
	}
	defer db.Close()
  
	err = db.Ping()
	rows, err := db.Query("select trx_id, trx_amt, to_char(trx_date, 'yyyymmdd'), trx_time, ref_num from reconciliation.t_source;")
	
	if err != nil {
		panic(err)
	}
	
	var (
		trx_id string
		trx_amt string
		trx_date string
		trx_time string
		ref_num string
	)
	
	for rows.Next() {
		if err = rows.Scan(&trx_id, &trx_amt, &trx_date, &trx_time, &ref_num); err != nil {
			fmt.Println("error fetching", err)
		}
		fmt.Println(trx_id, trx_amt, trx_date, trx_time, ref_num)
	}
	// fmt.Println("Successfully connected!")
  }