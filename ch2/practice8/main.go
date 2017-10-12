package main

import "fmt"

var vendorAddresses map[string]map[string]string = map[string]map[string]string{}

func main() {

	vendorAddress := map[string]string{"name": "", "vendor_addr": "", "vendor_addr_cn": "", "vendor_addr_tw": "", "vendor_geo": ""}
	vendorAddresses["danny"] = vendorAddress
	vendorAddress["name"] = "danny"
	vendorAddress["vendor_addr"] = "123456"

	fmt.Println(vendorAddresses)
}
