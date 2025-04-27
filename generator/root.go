package generator

import (
	"github.com/spf13/cobra"
	"os"
)

func Root(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		err := cmd.Usage()
		if err != nil {
			os.Exit(0)
		}
		return
	}

	parse(args)
}
