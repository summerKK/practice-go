package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println(a)

	ss := []string{"1_surface", "3_ipad", "5_mac pro", "7_mac air", "2_think pad", "6_idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)

	n := time.Now().Format("2006-01-02 15:04:05")

	fmt.Println(n)


	m := make(map[string]string,10)
	fmt.Println(len(m))

	fmt.Println()
}

//作者：t0n9
//链接：http: //www.jianshu.com/p/6e52bad56e06
//來源：简书
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
