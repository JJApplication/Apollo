/*
Project: dirichlet cmd_runner.go
Created: 2021/11/27 by Landers
*/

package utils

import (
	"context"
	"os/exec"
)

// 命令行运行
const (
	BASH = "bash"
	RUN  = "-c"
)

// 异步的命令行只适用于非及时返回的函数

func newCMD(envs []string, cmd ...string) *exec.Cmd {
	c := exec.Command(BASH, append([]string{RUN}, cmd...)...)
	c.Env = envs
	return c
}

func CMDAsync(envs []string, cmd ...string) error {
	return newCMD(envs, cmd...).Run()
}

func CMDRun(envs []string, cmd ...string) (string, error) {
	b, err := newCMD(envs, cmd...).Output()
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func CMDRaw(envs []string, cmd ...string) *exec.Cmd {
	return newCMD(envs, cmd...)
}

// CMDWithContext 执行结束后会杀死父进程来关闭shell
func CMDWithContext(envs []string, cmd ...string) (string, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	exe := exec.CommandContext(ctx, BASH, append([]string{RUN}, cmd...)...)
	exe.Env = envs
	b, err := exe.Output()
	if err != nil {
		return "", err
	}

	return string(b), err
}
