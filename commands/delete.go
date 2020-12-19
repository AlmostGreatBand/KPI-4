package commands

import (
	"github.com/AlmostGreatBand/KPI-4/eventloop"
)

type Delete struct {
	Args []string
}

func (d *Delete) Execute(h eventloop.Handler) {
	if len(d.Args) != 2 {
		h.Post(&Println{Message: "error: mismatch in quantity of arguments [split]"}, true)
		return
	}
	str := d.Args[0]
	symb := d.Args[1]

	res := ""
	for _, c := range str {
		char := string(c)
		if char != symb {
			res += char
		}
	}
	h.Post(&Println{ Message: res }, true)
}
