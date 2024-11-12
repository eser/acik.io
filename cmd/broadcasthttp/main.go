package main

import "github.com/eser/acik.io/pkg/broadcasthttp"

func main() {
	err := broadcasthttp.Run()
	if err != nil {
		panic(err)
	}
}
