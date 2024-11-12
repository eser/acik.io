package main

import "github.com/eser/acik.io/pkg/broadcastsvc"

func main() {
	err := broadcastsvc.Run()
	if err != nil {
		panic(err)
	}
}
