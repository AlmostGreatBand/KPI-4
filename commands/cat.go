package commands

import (
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strings"
)

type Cat struct {
	Args []string
}

func (c *Cat) Execute(h eventloop.Handler) {
	if len(c.Args) < 1 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [cat]" }, true)
		return
	}

	h.Post(&Println{ Message: strings.Join(c.Args, "") }, true)
}
