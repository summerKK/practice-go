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
	m := map[string]string{"name": "s'u'm'mer", "age": "25",}

	vals := make([]interface{}, 0, len(m))
	keys := []string{}
	prepare := []string{}
	for i, v := range m {
		vals = append(vals, v)
		fmt.Printf("%p\n\r", vals)
		keys = append(keys, i)
		prepare = append(prepare, "?")
	}
	fmt.Printf("%#v\n\r", vals)
	fmt.Println("----------")
	stmt, err := db.Prepare(`INSERT INTO user ( ` + strings.Join(keys, ",") + `) VALUES ( ` + strings.Join(prepare, ",") + `)`)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(`INSERT INTO user ( ` + strings.Join(keys, ",") + `) VALUES ( ` + strings.Join(prepare, ",") + `)`)
	fmt.Println(vals)
	rows, err := stmt.Exec(vals...)

	defer stmt.Close()

	//rows.Close()
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}

	fmt.Println(rows.LastInsertId())
}

func selectData(db *sql.DB) {
	var id int
	var name string
	var age string
	m := map[string]string{"summer":"10"}
	stmt, _ := db.Prepare(`SELECT * From user where age > ?`)

	rows, err := stmt.Query(m["summer"])

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

	//deleteData(db)

	insertData(db)

	//selectData(db)
}
