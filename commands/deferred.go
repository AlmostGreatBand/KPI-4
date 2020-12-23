package commands

import (
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"strconv"
	"strings"
	"time"
)

type Deferred struct {
	Args []string
}

func (d *Deferred) Execute(h eventloop.Handler) {
	if len(d.Args) < 2 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [deferred]" }, true)
		return
	}

	t, err := strconv.Atoi(d.Args[0])
	if err != nil {
		h.Post(&Println{ Message: "error: param is not a number [deferred]" }, true)
		return
	}

	message := strings.Join(d.Args[1:], " ")


	go func() {
		time.Sleep(time.Duration(t) * time.Second)
		h.Post(&Println{ Message: message }, true)
	}()
}