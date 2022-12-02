package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	"github.com/clemensgg/flash/app"
	"github.com/clemensgg/flash/cmd"
)

func main() {
	rootCmd, _ := cmd.NewRootCmd(
		app.AppName,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		app.DefaultChainID,
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)

	rootCmd.AddCommand(cmd.AddConsumerSectionCmd(app.DefaultNodeHome))

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
