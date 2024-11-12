package main

import "github.com/eser/acik.io/pkg/testhttp"

func main() {
	err := testhttp.Run()
	if err != nil {
		panic(err)
	}
}
