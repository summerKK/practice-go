package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strings"
	"sync"
)

var (
	connection *sql.DB
)

func main() {
	db, err := sql.Open("mysql", "root:177678012@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	connection = db

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		colums := make([]interface{}, 3)
		columsPtr := make([]interface{}, 3)

		for i := range colums {
			columsPtr[i] = &colums[i]
		}

		rows.Scan(columsPtr...)
		fmt.Println(string(colums[1].([]byte)))
	}

	rows.Close()

	m := []map[string]string{
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
		map[string]string{"name": "s'u'm'mer", "age": "25",},
	}
	var wg sync.WaitGroup
	for _, v := range m {
		wg.Add(1)
		go func(m map[string]string) {
			insertData(m)
			wg.Done()
		}(v)
	}

	wg.Wait()
	fmt.Println("Done")

	var in interface{} = 1

	v, ok := in.([]byte)
	if ok {
		fmt.Println(v, "convert")
	} else {
		fmt.Println(in, "no")
	}

	i := fmt.Sprintf("%s", 12)
	fmt.Println(i)

	var mm []string

	mm = append(mm, "summer")

	fmt.Println(mm)

	fmt.Println("hello")

	var lock sync.Mutex
	lock.Lock()
	fmt.Println("hello")
	lock.Unlock()

	tmp := []string{"hello", "world"}
	dst := map[string][]string{}
	dst["summer"] = tmp
	tmp = []string{"11", "22"}
	dst["danny"] = tmp

	fmt.Println(dst)

	ss := make([]string,0,10)
	ss = []string{"1","2"}
	fmt.Println(cap(ss))

	ss = []string{}
	fmt.Println(cap(ss))
}

func insertData(m map[string]string) {

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
	stmt, err := connection.Prepare(`INSERT INTO user ( ` + strings.Join(keys, ",") + `) VALUES ( ` + strings.Join(prepare, ",") + `)`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(`INSERT INTO user ( ` + strings.Join(keys, ",") + `) VALUES ( ` + strings.Join(prepare, ",") + `)`)
	fmt.Println(vals)
	rows, err := stmt.Query(vals...)

	defer stmt.Close()

	rows.Close()
	if err != nil {
		log.Fatalf("insert data error: %v\n", err)
	}
}
