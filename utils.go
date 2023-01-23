package main

import (
	"bytes"
	"os"
	"os/exec"
)

func FileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func DirectoryExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

func CmdExec(args ...string) (*bytes.Buffer, error) {
	baseCmd := args[0]
	cmdArgs := args[1:]

	cmd := exec.Command(baseCmd, cmdArgs...)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	cmd.Stderr = cmdOutput
	err := cmd.Run()
	if err != nil {
		return cmdOutput, err
	}
	return cmdOutput, nil
}

func CmdExecConsole(args ...string) error {
	baseCmd := args[0]
	cmdArgs := args[1:]
	cmd := exec.Command(baseCmd, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func CmdExecStrOutput(args ...string) (string, error) {
	res, err := CmdExec(args...)
	return res.String(), err
}
