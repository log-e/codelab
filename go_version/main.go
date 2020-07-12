package main

import (
	"flag"
	"fmt"
	v "go_version/version"
	"os"
)

var version bool

func init(){
	flag.BoolVar(&version, "v", false, "version info")
	flag.Parse()
}

func main() {
	if version {
		fmt.Println(v.Get())
		os.Exit(0)
	}
	fmt.Println("Running")
}