package main

import (
	"fmt"
	"io"

	"github.com/fatih/color"
	"github.com/gosuri/uitable"
	"github.com/tj/go-spin"
)

var (
	loggerUI = color.New(color.FgCyan, color.Bold).SprintFunc()
	msgUI    = color.New(color.FgWhite, color.Bold).SprintFunc()
	labelUI  = color.New(color.FgCyan, color.Bold).SprintFunc()
	spinner  = spin.New()
)

type Style struct{}

func init() {
	spinner.Set(spin.Spin6)
}

func displayMap(writer io.Writer, props map[string]interface{}) {
	table := uitable.New()
	for key, value := range props {
		table.AddRow("", labelUI(key), value)
	}
	fmt.Fprintln(writer, table)
	fmt.Fprintln(writer)
}

// TODO support for smart colors (depending on types, log levels, namespaces, ...)
func display(writer io.Writer, logData StructuredLog, unfold bool) {
	fmt.Fprintf(writer, "\n\t%s\t( %d )\tprocessing logs ...\n\n", spinner.Next(), counter)

	fmt.Fprintf(writer, "\t%-30s | %s\n\n", loggerUI(logData.Logger), msgUI(logData.Msg))
	if unfold {
		displayMap(writer, logData.Properties)
	}
}
