package subdir

var (
	Ch     chan int
	Symbol bool = false
)

func Start() {
	Ch = make(chan int)
	Ch <- 1
	Symbol = true
}
