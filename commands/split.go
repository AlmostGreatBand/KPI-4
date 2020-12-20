package commands

import (
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strings"
)

type Split struct {
	Args []string
}

func (s *Split) Execute(h eventloop.Handler) {
	if len(s.Args) != 2 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [split]" }, true)
		return
	}
	str := s.Args[0]
	sep := s.Args[1]

	values := strings.Split(str, sep)
	for _, value := range values {
		h.Post(&Println{ Message: value }, true)
	}
}