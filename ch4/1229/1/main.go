package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"strings"
	"log"
)

func main() {
	attrJson := map[string]string{}
	attr := `{"Brand":"Hermès","Condition":"New or Never Worn","Designer":"Hermès","Dimensions":"24 cm Hx30 cm Wx15 cm D","Material Notes":"Leather","Period":"21st Century","Place of Origin":"France","Seller Location":"New York, NY"}`
	err := json.Unmarshal([]byte(attr), &attrJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	attrJson1 := map[string]string{}
	attr1 := ``
	json.Unmarshal([]byte(attr1), &attrJson1)
	fmt.Println(attrJson)
	fmt.Println(attrJson1)
	fmt.Println(len(attrJson1))

	attrs := []map[string]string{}
	var wg sync.WaitGroup
	var mux sync.Mutex
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mux.Lock()
			attrs = append(attrs, map[string]string{"summer": "hello"})
			mux.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(attrs)

	type summer struct {
		arr1 []map[string]string
		arr2 map[string]string
	}

	var struct1 summer
	struct1.arr1 = append(struct1.arr1, map[string]string{"summer": "hello"})

	s := `UPDATE lux_products SET market_price = ? , special_price = ? , `
	fmt.Println(strings.TrimSuffix(s,", "))

	log.Println(strings.TrimRight("abba", "ba"))
	log.Println(strings.TrimRight("abcdaaaaa", "abcd"))
	log.Println(strings.TrimSuffix("abcddcba", "dcba"))

}
