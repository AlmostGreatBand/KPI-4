package commands

import (
	"fmt"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
)

type Println struct {
	Message string
}

func (p *Println) Execute(_ eventloop.Handler) {
	fmt.Println(p.Message)
}