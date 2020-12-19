package parser

import (
	"bufio"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func Parse() {
	file, err := os.OpenFile("example.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}
	defer file.Close()

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

		_, _ = parseSha1Expression(string(line))
	}
}

func parseSha1Expression(str string) (string, error) {
	values := strings.Split(str, " ")

	if values[0] != "sha1"  {
		return "", errors.New("error occurred while parsing expression")
	}
	for _, value := range values[1:] {
		encryptor := sha1.New()
		encryptor.Write([]byte(value))

		res := encryptor.Sum(nil)

		fmt.Printf("%s - %x\n", value,  res)
	}
	return "", nil
}