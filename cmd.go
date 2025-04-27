package framework

import (
	"fmt"
	"github.com/jackylany/framework/generator"
	"github.com/spf13/cobra"
	"os"
)

var (
	h       bool
	rootCmd = &cobra.Command{
		Use:   "Sita",
		Short: "Sita for generate api framework",
		Long:  `Quick to make api project`,
		Run: func(cmd *cobra.Command, args []string) {
			generator.Root(cmd, args)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVar(&h, "h", false, "")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
