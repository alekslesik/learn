package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/alekslesik/config"
	_ "github.com/go-sql-driver/mysql"
)


func main() {
	db, err := sql.Open("mysql", "root:486464@/kinovdom")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT ID FROM b_catalog_product")
	if err != nil {
		log.Fatalln(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id)
	}

	err = rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	config.New()

}

