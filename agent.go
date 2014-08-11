package gatekeeper

import (
	"strconv"

	"github.com/mindfork/mindfork"
)

type Gatekeeper struct {
	mindfork.Agent
}

func (gk *Gatekeeper) Name() string {
	return "gatekeeper/" + strconv.Itoa(gk.Id())
}
