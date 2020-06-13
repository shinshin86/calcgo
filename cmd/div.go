package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func divCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "div",
		Short: "division culc",
		Args:  cobra.RangeArgs(2, 2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := divAction(args); err != nil {
				Exit(err, 1)
			}
		},
	}

	return cmd
}

func divAction(args []string)(err error) {
	n1, err := strconv.Atoi(args[0])

	if err != nil {
		return err
	}

	n2, err := strconv.Atoi(args[1])

	if err != nil {
		return err
	}

	fmt.Println(string(ColorGreen), n1 / n2)

	return nil
}
