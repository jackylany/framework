package framework

import (
	"fmt"
	"github.com/jackylany/framework/common"
	"github.com/jackylany/framework/generator"
	"github.com/jackylany/framework/generator/gen"
	"github.com/spf13/cobra"
	"os"
)

var (
	h        bool
	empty    bool
	fileName string
	tempPath string
)

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "Tool for generate api framework",
	Long:  `Quick to make api project`,
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "version",
	Aliases: []string{"v"},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(common.Version)
	},
}

var controllerCmd = &cobra.Command{
	Use:     generator.OpMakeController,
	Short:   "make a controller",
	Long:    `make a controller to api project`,
	Example: `make:controller user [--empty]`,
	Run: func(cmd *cobra.Command, args []string) {
		if empty {
			tempPath = fmt.Sprintf("%s/../../temp/make/empty", gen.FrameworkRoot())
		}
		fmt.Printf("%s %s\n", fileName, tempPath)
		if len(args) < 1 {
			fmt.Println("missing controller name")
			return
		}
		generator.Make(&gen.Controller{Name: args[0], TempPath: tempPath})
	},
}

var modelCmd = &cobra.Command{
	Use:   generator.OpMakeModel,
	Short: "make a model",
	Long:  `make a model to api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("missing model name")
			return
		}

		fileName = args[1]
		generator.Make(&gen.Model{Name: fileName, TempPath: tempPath})
	},
}

var middlewareCmd = &cobra.Command{
	Use:   generator.OpMakeMiddleware,
	Short: "make a middleware",
	Long:  `make a middleware to api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("missing middleware name")
			return
		}

		fileName = args[1]
		generator.Make(&gen.Middleware{Name: fileName, TempPath: tempPath})
	},
}

var transformerCmd = &cobra.Command{
	Use:   generator.OpMakeTransformer,
	Short: "make a transformer",
	Long:  `make a transformer to api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("missing transformer name")
			return
		}

		fileName = args[1]
		generator.Make(&gen.Transform{Name: fileName, TempPath: tempPath})
	},
}

var routeCmd = &cobra.Command{
	Use:   generator.OpMakeRoute,
	Short: "make a route",
	Long:  `make a route to api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("missing route name")
			return
		}

		fileName = args[1]
		generator.Make(&gen.Route{Name: fileName, TempPath: tempPath})
	},
}

var validateCmd = &cobra.Command{
	Use:   generator.OpMakeValidate,
	Short: "make a validate",
	Long:  `make a validate to api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("missing validate name")
			return
		}

		fileName = args[1]
		generator.Make(&gen.Validator{Name: fileName, TempPath: tempPath})
	},
}

var parameterCmd = &cobra.Command{
	Use:   generator.OpMakeParameter,
	Short: "make a parameter",
	Long:  `make a parameter to api project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("missing parameter name")
			return
		}

		fileName = args[1]
		generator.Make(&gen.Validator{Name: fileName, TempPath: tempPath})
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&h, "h", false, "")
	controllerCmd.PersistentFlags().BoolVar(&empty, "empty", false, "make a empty controller")
	rootCmd.AddCommand(
		versionCmd,
		controllerCmd,
		modelCmd,
		middlewareCmd,
		transformerCmd,
		routeCmd,
		validateCmd,
		parameterCmd,
	)
	tempPath = fmt.Sprintf("%s/../../temp/make", gen.FrameworkRoot())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
