package command

import (
	"fmt"
	"os"
	"path"

	"github.com/pefish/pefish-tool/pkg/global"
	"github.com/pkg/errors"

	"github.com/pefish/go-commander"
	go_file "github.com/pefish/go-file"
	go_shell "github.com/pefish/go-shell"
)

type InstallBinCommand struct {
}

func NewInstallBinCommand() *InstallBinCommand {
	return &InstallBinCommand{}
}

func (dc *InstallBinCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *InstallBinCommand) Data() interface{} {
	return nil
}

func (dc *InstallBinCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *InstallBinCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *InstallBinCommand) Start(command *commander.Commander) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	targetPath := path.Join(homeDir, "pefish-bins")
	exist, err := go_file.Exists(targetPath)
	if err != nil {
		return err
	}
	if exist {
		return errors.Errorf("<%s> 目录已经存在。", targetPath)
	}
	cmd := go_shell.NewCmd("git clone https://github.com/pefish/pefish-bins.git %s", targetPath)
	err = go_shell.ExecInConsole(cmd)
	if err != nil {
		return err
	}

	exportPathStr := fmt.Sprintf(`export PATH="$PATH:%s"`, targetPath)

	err = go_shell.ExecInConsole(go_shell.NewCmd(`sudo echo '%s' >> /etc/profile`, exportPathStr))
	if err != nil {
		return err
	}

	command.Logger.Info("Done.")

	return nil
}
