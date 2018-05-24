package main

import (
	"fmt"
	"runtime"
)

func say(s string) {

	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)

	}
}

func main() {

	//fmt.Println("Hello ")
	runtime.Gosched()

	go say("world")

	say("Hello,")

	// go say("world")
}
