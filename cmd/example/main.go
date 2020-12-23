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
	"time"
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

	/*
		main doesn't wait for child goroutines,
		need it here just to show how deferred call works
	 */
	time.Sleep(10 * time.Second)
}

func parse(line string) eventloop.Command {
	values := strings.Fields(strings.TrimSpace(line))
	if len(values) < 1 {
		return &commands.Println{ Message: "error: blank line" }
	}

	command := values[0]
	args := values[1:]

	switch command {
	case "sha1": return &commands.Sha1{ Args: args }
	case "print": return &commands.Print{ Message: strings.Join(args, " ") }
	case "println": return &commands.Println{ Message: strings.Join(args, " ") }
	case "printc": return &commands.Printc{ Args: args }
	case "add": return &commands.Add{ Args: args }
	case "reverse": return &commands.Reverse{ Args: args }
	case "palindrome": return &commands.Palindrome{ Args: args }
	case "split": return &commands.Split{ Args: args }
	case "delete": return &commands.Delete{ Args: args }
	case "cat": return &commands.Cat{ Args: args }
	case "multiply": return &commands.Multiply{ Args: args }
	case "deferred": return &commands.Deferred{ Args: args }
	default: return &commands.Println{ Message: "error: unexpected method name" }
	}
}
