package main

import (
	"flag"
	"time"
	"fmt"
	"sort"
)

type tm1 struct {
	m int
}

type tm2 struct {
	tm1
}

type stringSlice []string

var period = flag.Duration("period", 1*time.Second, "time period")

func main() {

	flag.Parse()
	fmt.Println(fmt.Sprintf("sleeping %v", *period))
	time.Sleep(*period)
	fmt.Println()

	s := "12.00ss"
	var unit float64
	var value string
	fmt.Println(fmt.Sscanf(s, "%f%s", &unit, &value))
	fmt.Println(unit, value)

	var c *tm2 = &tm2{}
	c.tm1.m = 2
	c.E()
	fmt.Println(c.tm1.m)

	var d stringSlice = stringSlice{"b","a","1","c"}
	sort.Sort(d)
	fmt.Println(d)


}

func (c *tm1) E() {
	fmt.Println("Hello World")
	fmt.Println(c.m)
	c.m = 5
}

func (p stringSlice) Len() int {
	return len(p)
}

func (p stringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p stringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
