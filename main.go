package main

import (
	"context"

	"github.com/edaniels/golog"
	"go.viam.com/rdk/components/base"
	"go.viam.com/rdk/module"
	"go.viam.com/utils"
)

func main() {
	utils.ContextualMain(mainWithArgs, golog.NewDevelopmentLogger("limo"))
}

func mainWithArgs(ctx context.Context, args []string, logger golog.Logger) error {
	limoBaseModule, err := module.NewModuleFromArgs(ctx, logger)
	if err != nil {
		return err
	}

	limoBaseModule.AddModelFromRegistry(ctx, base.API, Model)

	err = limoBaseModule.Start(ctx)
	defer limoBaseModule.Close(ctx)
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}
