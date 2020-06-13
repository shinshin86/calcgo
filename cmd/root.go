package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "culc",
		Short: "command line calculator",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("room command")
		},
	}
)

func Run() {
	rootCmd.Execute()
}

func Exit(err error, codes ...int) {
	var code int
	if len(codes) > 0 {
		code = codes[0]
	} else {
		code = 2
	}
	if err != nil {
		fmt.Println(string(ColorRed), err)
	}
	os.Exit(code)
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(
		sumCommand(),
		subCommand(),
		multCommand(),
		divCommand(),
	)
}
