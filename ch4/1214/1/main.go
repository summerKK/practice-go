package main

import "fmt"

func main() {
	p := map[string]string{}
	fmt.Printf("%p\n\r", &p)
	test1(p)
	fmt.Println(p) // map[summer:summer]

	var p2 map[string]string
	test2(p2)
	fmt.Println(p2) // map[]

	p3 := map[string]string{}
	test3(p3)
	fmt.Println(p3) // map[]

	p4 := make([]int, 2)
	test4(p4)
	fmt.Println(p4, len(p4)) // [0,0] 2

	fmt.Println("------------p5---------------")
	p5 := map[string]string{}
	fmt.Printf("%p\n\r", p5)
	test5(&p5)
	fmt.Println(p5)

	fmt.Println("------------p6---------------")
	var p6 map[string]string
	fmt.Printf("%p\n\r", p6)
	test6(&p6)
	fmt.Println(p6)

	fmt.Println("------------p7---------------")
	p7 := make([]int, 2)
	test7(&p7)
	fmt.Printf("%p\n\r", p7)
	fmt.Println(p7) // [0 0 1 2]

	fmt.Println("------------p8---------------")
	p8 := make([]int,3,4)
	test8(p8)
	fmt.Println(p8)

}

func test1(p map[string]string) {
	fmt.Printf("%p\n\r", &p)
	p["summer"] = "summer"
}

func test2(p map[string]string) {
	p = make(map[string]string)
	p["hello"] = "world"
}

func test3(p map[string]string) {
	p = map[string]string{
		"hello": "world",
	}
}

func test4(p []int) {
	p = append(p, 1)
	p = append(p, 2)
	fmt.Println(p, len(p)) // [0 0 1 2 ] 4
}

func test5(p *map[string]string) {
	fmt.Printf("%p\n\r", *p)
}

func test6(p *map[string]string) {
	fmt.Printf("%p\n\r", p)
	fmt.Printf("%p\n\r", &p)
	fmt.Printf("%p\n\r", *p)
}

func test7(p *[]int) {
	*p = append(*p, 1)
	*p = append(*p, 2)
	fmt.Println(&*p)
	fmt.Printf("%p\n\r", *p)
}

func test8(p []int)  {
	p[0] = 1
	fmt.Printf("%p\n",p)
	p = append(p,2)
	fmt.Printf("%p\n",p)
	p[2] = 4
	fmt.Println(p)
}
