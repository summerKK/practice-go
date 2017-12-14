package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]string{
		"1": "2",
		"2": "2",
		"3": "2",
		"4": "2",
	}

	enc, _ := json.Marshal(m)
	fmt.Println(string(enc))
}
