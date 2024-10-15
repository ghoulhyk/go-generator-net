package main

import (
	"github.com/ghoulhyk/go-generator-net/cmd/generate"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{Use: "netGenerator"}
	cmd.AddCommand(
		generate.Cmd(),
	)
	_ = cmd.Execute()
}
