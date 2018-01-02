package main

import (
	"encoding/json"
	"fmt"
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
}
