package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"

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

// TODO Regex
// FIXME crash on bad filters without --strict
func matched(logData *StructuredLog, filters map[string]string, strict bool) bool {
	for k, v := range filters {
		// check msg
		if k == "msg" && logData.Msg == v {
			return true
		}
		// check logger
		if k == "logger" && logData.Logger == v {
			return true
		}
		// check properties
		pvalue, ok := logData.Properties[k]
		if strict && !ok {
			return false
		}
		match, _ := regexp.MatchString(v, pvalue.(string))
		if !match {
			return false
		}
	}
	return true
}

func loop() int {
	wasMatched := false
	opts := getopt()
	writer := uilive.New()
	// start listening for updates and render
	writer.Start()
	defer writer.Stop()

	for {
		logData, err := stdingSerialize()
		if err == io.EOF {
			log.Printf("reached EOF, exiting")
			return 0
		} else if err != nil {
			log.Printf("failed to parse JSON: %s", err)
			return 1
		}

		counter++
		if wasMatched {
			wasMatched = false
			display(os.Stdout, *logData, opts.Unfold)
		} else {
			display(writer, *logData, opts.Unfold)
		}
		if matched(logData, opts.Filters, opts.Strict) {
			wasMatched = true
		}
	}
}

func main() {
	os.Exit(loop())
}
