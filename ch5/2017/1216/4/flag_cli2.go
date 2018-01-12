package main

import (
	"github.com/jessevdk/go-flags"
	"fmt"
)

var opts struct {
	Name     string `short:"n" long:"name" default:"world" description:"A name to say hello."`
	Spainish bool   `short:"s" long:"spanish" description:"use spanish language"`
}

func main() {
	flags.Parse(&opts)

	if opts.Spainish == true {
		fmt.Printf("Hole %s !\n", opts.Name)
	}else {
		fmt.Printf("Hello %s !\n",opts.Name)
	}
}
