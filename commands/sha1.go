package commands

import (
	"crypto/sha1"
	"fmt"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
)

type Sha1 struct {
	Args []string
}

func (s *Sha1) Execute(h eventloop.Handler) {
	if len(s.Args) < 1 {
		h.Post(&Println{ Message: "error: mismatch in quantity of arguments [sha1]" }, true)
		return
	}

	for _, value := range s.Args {
		encryptor := sha1.New()
		encryptor.Write([]byte(value))

		res := encryptor.Sum(nil)
		resHex := fmt.Sprintf("%x", res)

		h.Post(&Println{ Message: resHex }, true)
	}
}
