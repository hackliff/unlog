package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
)

// TODO version
type Options struct {
	Version bool `short:"v" long:"version" description:"print version and exit"`
	Strict bool `short:"s" long:"strict" description:"discard logs lacking filter keys"`
	Unfold  bool `short:"u" long:"unfold" description:"also dump log properties"`
	Filters map[string]string `short:"f" long:"filter" description:"filter logs to display based on their embedded properties"`
}

func getopt() *Options {
	opts := &Options{}
	if _, err := flags.Parse(opts); err != nil {
		// NOTE will also exit on '--help' (exit code 1 is rude)
		os.Exit(1)
	}

	if opts.Version {
		fmt.Printf("%s v%s\n", "unlog", VERSION)
		os.Exit(0)
	}

	return opts
}
