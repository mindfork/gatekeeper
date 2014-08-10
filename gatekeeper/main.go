package main

import "github.com/mindfork/gatekeeper"

func main() {
	gk := &gatekeeper.Gatekeeper{}

	gk.Init(":25001")
	gk.Run()
}
