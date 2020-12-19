package commands

import (
	"fmt"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strconv"
	"strings"
)

type Printc struct {
	Args []string
}

func (p *Printc) Execute(h eventloop.Handler) {
	if len(p.Args) != 2 {
		h.Post(&Println{ Message: "error: missmatch in quantity of arguments [printc]" }, true)
	}
	str := p.Args[0]

	if count, err := strconv.Atoi(p.Args[1]); err != nil {
		h.Post(&Println{ Message: fmt.Sprintf("error: second parem is not a number [printc] - %s", err) }, true)
	} else {
		res := strings.Repeat(str, count)
		h.Post(&Println{ Message: res }, true)
	}

}