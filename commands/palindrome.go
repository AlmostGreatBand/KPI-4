package commands

import (
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strings"
)

type Palindrome struct {
	Args []string
}

func (p *Palindrome) Execute(h eventloop.Handler) {
	length := len(p.Args)
	if length < 1 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [palindrome]" }, true)
		return
	}

	res := make([]string, length)
	for i, word := range p.Args {
		reversed := ""
		for _, char := range word {
			reversed = string(char) + reversed
		}
		res[i] = word + reversed
	}


	h.Post(&Println{ Message: strings.Join(res, " ") }, true)
}

