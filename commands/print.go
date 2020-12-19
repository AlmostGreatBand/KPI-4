package commands

import (
	"fmt"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
)

type Print struct {
	Message string
}

func (p *Print) Execute(_ eventloop.Handler) {
	fmt.Print(p.Message)
}
