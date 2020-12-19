package commands

import (
	"fmt"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strconv"
)

type Add struct {
	Args []string
}

func (a *Add) Execute(h eventloop.Handler) {
	if len(a.Args) < 1 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [printc]" }, true)
	}

	res := 0
	for _, i := range a.Args {
		if value, err := strconv.Atoi(i); err != nil {
			h.Post(&Println{ Message: fmt.Sprintf("error: parem is not a number [printc] - %s", err) }, true)
		} else {
			res += value
		}
	}

	h.Post(&Println{ Message: strconv.Itoa(res) }, true)
}
