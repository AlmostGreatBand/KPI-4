package commands

import (
	"fmt"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strconv"
)

type Multiply struct {
	Args []string
}

func (m *Multiply) Execute(h eventloop.Handler) {
	if len(m.Args) < 1 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [add]" }, true)
		return
	}

	res := 1
	for _, i := range m.Args {
		if value, err := strconv.Atoi(i); err != nil {
			h.Post(&Println{ Message: fmt.Sprintf("error: parem is not a number [add] - %s", err) }, true)
		} else {
			res *= value
		}
	}

	h.Post(&Println{ Message: strconv.Itoa(res) }, true)
}
