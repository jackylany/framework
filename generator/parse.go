package generator

import (
	"fmt"
	"github.com/jackylany/framework/generator/gen"
)

const (
	OpMakeController  = "make:controller"
	OpMakeModel       = "make:model"
	OpMakeTransformer = "make:transformer"
	OpMakeRoute       = "make:route"
	OpMakeMiddleware  = "make:middleware"
	OpMakeMigrate     = "make:migrate"
	OpMakeValidate    = "make:validate"
	OpMakeParameter   = "make:parameter"
	OpMakeModule      = "make:module"
	OpMigrate         = "migrate:run"
)

func parse(args []string) error {

	op := args[0]

	fileName := args[1]

	tempPath := fmt.Sprintf("%s/../../temp/make", gen.FrameworkRoot())

	switch op {
	case OpMakeController:
		Make(&gen.Controller{Name: fileName, TempPath: tempPath})
	case OpMakeModel:
		Make(&gen.Model{Name: fileName, TempPath: tempPath})
	case OpMakeTransformer:
		Make(&gen.Transform{Name: fileName, TempPath: tempPath})
	case OpMakeRoute:
		Make(&gen.Route{Name: fileName, TempPath: tempPath})
	case OpMakeMiddleware:
		Make(&gen.Middleware{Name: fileName, TempPath: tempPath})
	case OpMakeParameter:
		Make(&gen.Parameter{Name: fileName, TempPath: tempPath})
	case OpMakeValidate:
		Make(&gen.Validator{Name: fileName, TempPath: tempPath})
		/*case OpMakeMigrate:
		Make(&gen.Transform{Name: fileName})
		*/
	}

	return nil
}
