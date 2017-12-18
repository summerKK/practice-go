package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strings"
)

func deleteData(db *sql.DB) {
	stmt, _ := db.Prepare(`DELETE FROM user WHERE id = ?`)

	rows, err := stmt.Query(1)
	defer stmt.Close()

	rows.Close()
	if err != nil {
		log.Fatalf("delete data error: %v\n", err)
	}

	rows, err = stmt.Query(2)
	rows.Close()
	if err != nil {
		log.Fatalf("delete data error: %v\n", err)
	}
}

func insertData(db *sql.DB) {
	m := map[string]string{"name": "s'u'm'mer", "age": "25"}

	vals := []interface{}{}
	keys := []string{}
	prepare := []string{}
	for i, v := range m {
		vals = append(vals, v)
		keys = append(keys, i)
		prepare = append(prepare, "?")
	}

	stmt, _ := db.Prepare(`INSERT INTO user ( ` + strings.Join(keys, ",") + `) VALUES ( ` + strings.Join(prepare, ",") + `)`)
	fmt.Println(`INSERT INTO user ( ` + strings.Join(keys, ",") + `) VALUES ( ` + strings.Join(prepare, ",") + `)`)
	fmt.Println(vals)
	rows, err := stmt.Query(vals...)

	defer stmt.Close()

	rows.Close()
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}
}

func selectData(db *sql.DB) {
	var id int
	var name string
	var age int
	stmt, _ := db.Prepare(`SELECT * From user where age > ?`)

	rows, err := stmt.Query(10)

	defer stmt.Close()
	defer rows.Close()

	if err != nil {
		log.Fatalf("select data error: %v\n", err)
		return
	}
	for rows.Next() {
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("get data, id: %d, name: %s, age: %d", id, name, age)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	db, err := sql.Open("mysql", "root:177678012@tcp(127.0.0.1:3306)/test")

	defer db.Close()

	if err != nil {
		fmt.Printf("connect to db 127.0.0.1:3306 error: %v\n", err)
		return
	}

	deleteData(db)

	insertData(db)

	selectData(db)
}
