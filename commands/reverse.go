package commands

import (
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strings"
)

type Reverse struct {
	Args []string
}

func (r *Reverse) Execute(h eventloop.Handler) {
	length := len(r.Args)
	if length < 1 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [reverse]" }, true)
		return
	}

	res := make([]string, length)
	for i, word := range r.Args {
		reversed := ""
		for _, char := range word {
			reversed = string(char) + reversed
		}
		res[i] = reversed
	}

	h.Post(&Println{ Message: strings.Join(res, " ") }, true)
}
