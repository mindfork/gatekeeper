package gatekeeper

import "github.com/mindfork/mindfork"

type Gatekeeper struct {
	mindfork.Agent
	name string
}

func (gk *Gatekeeper) Init(port string) {
	gk.Agent.Init(port)
	gk.name = "Gatekeeper"
}
