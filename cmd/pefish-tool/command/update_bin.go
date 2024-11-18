package command

import (
	"os"
	"path"

	"github.com/pefish/pefish-tool/pkg/global"
	"github.com/pkg/errors"

	"github.com/pefish/go-commander"
	go_file "github.com/pefish/go-file"
	go_shell "github.com/pefish/go-shell"
)

type UpdateBinCommand struct {
}

func NewUpdateBinCommand() *UpdateBinCommand {
	return &UpdateBinCommand{}
}

func (dc *UpdateBinCommand) Config() interface{} {
	return &global.GlobalConfig
}

func (dc *UpdateBinCommand) Data() interface{} {
	return nil
}

func (dc *UpdateBinCommand) Init(command *commander.Commander) error {
	return nil
}

func (dc *UpdateBinCommand) OnExited(command *commander.Commander) error {
	return nil
}

func (dc *UpdateBinCommand) Start(command *commander.Commander) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	targetPath := path.Join(homeDir, "pefish-bins")
	exist, err := go_file.Exists(targetPath)
	if err != nil {
		return err
	}
	if !exist {
		return errors.Errorf("<%s> 目录不存在。", targetPath)
	}
	cmd := go_shell.NewCmd("cd %s && git pull", targetPath)
	err = go_shell.ExecInConsole(cmd)
	if err != nil {
		return err
	}

	command.Logger.Info("Done.")

	return nil
}
