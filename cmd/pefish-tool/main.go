package main

import (
	"fmt"

	"github.com/pefish/pefish-tool/cmd/pefish-tool/command"
	"github.com/pefish/pefish-tool/version"

	"github.com/pefish/go-commander"
)

func main() {
	commanderInstance := commander.NewCommander(
		version.AppName,
		version.Version,
		fmt.Sprintf("%s is a private tool for pefish. Author: pefish", version.AppName),
	)
	commanderInstance.RegisterSubcommand("install-bin", &commander.SubcommandInfo{
		Desc:       "Use this command to install bins.",
		Args:       nil,
		Subcommand: command.NewInstallBinCommand(),
	})
	commanderInstance.RegisterSubcommand("update-bin", &commander.SubcommandInfo{
		Desc:       "Use this command to update bins.",
		Args:       nil,
		Subcommand: command.NewUpdateBinCommand(),
	})
	err := commanderInstance.Run()
	if err != nil {
		commanderInstance.Logger.Error(err)
	}
}
