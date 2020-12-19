package main

import (
	"bufio"
	"flag"
	"github.com/AlmostGreatBand/KPI-4/commands"
	"github.com/AlmostGreatBand/KPI-4/eventloop"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("f", "cmd/example/example.txt", "File with commands")
	flag.Parse()

	file, err := os.OpenFile(*path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer file.Close()

	var eventLoop eventloop.EventLoop
	eventLoop.Start()

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("read line error: %v", err)
			return
		}

		command := parse(string(line))
		eventLoop.Post(command, false)
	}
	eventLoop.AwaitFinish()
}

func parse(line string) eventloop.Command {
	values := strings.Split(line, " ")
	command := values[0]
	args := values[1:]

	switch command {
	case "sha1": return &commands.Sha1{ Args: args }
	case "print": return &commands.Print{ Message: strings.Join(args, " ") }
	case "println": return &commands.Println{ Message: strings.Join(args, " ") }
	case "printc": return &commands.Printc{ Args: args }
	case "add": return &commands.Add{ Args: args }
	default: return &commands.Print{ Message: "error: unexpected method name"}
	}
}
