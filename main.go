package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/gosuri/uilive"
)

const VERSION string = "0.2.0"

var (
	stdin   *bufio.Reader
	counter = 0
)

func init() {
	stdin = bufio.NewReader(os.Stdin)
}

func stdingSerialize() (*StructuredLog, error) {
	input, err := stdin.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	return parseLog(input)
}

func loop() int {
	opts := getopt()
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()

	for {
		logData, err := stdingSerialize()
		if err == io.EOF {
			writer.Stop()
			log.Printf("reached EOF, exiting")
			return 0
		} else if err != nil {
			log.Printf("failed to parse JSON: %s", err)
			return 1
		}

		counter++
		display(writer, *logData, opts.Unfold)
	}
}

func main() {
	os.Exit(loop())
}
