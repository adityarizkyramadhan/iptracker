package main

import (
	"fmt"
	"github.com/adityarizkyramadhan/iptracker"
)

func main() {
	data, err := iptracker.Trace("103.121.18.64")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(data)
}
