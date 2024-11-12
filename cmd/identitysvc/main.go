package main

import "github.com/eser/acik.io/pkg/identitysvc"

func main() {
	err := identitysvc.Run()
	if err != nil {
		panic(err)
	}
}
