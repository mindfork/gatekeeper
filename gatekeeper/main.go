package main

import "github.com/mindfork/gatekeeper"

func main() {
	gk := &gatekeeper.Gatekeeper{}

	gk.SetPort(":25001")
	gk.Init()
	gk.Run()
}
