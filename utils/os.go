package utils

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// 检查目录权限
func CheckDirMode() error {
	// 获取当前目录
	dir, err := os.Getwd()
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	// 目录是否可读可写
	return syscall.Access(dir, syscall.O_RDWR)

}

const ShellToUse = "bash"

func ExecShellCommand(command string) error {
	cmd := exec.Command(ShellToUse, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	return err
}

// 替换项目中的主包名
func ReplaceMainPackageNAme(pwd, name string) error {
	CommandLogger.Warning("Init", fmt.Sprintf("You create a new project,please replace main package name from 'goapp' to '%s'", name))
	// TODO 遍历项目每一个文件
	// TODO 读取每个文件前n行
	// TODO 替换成特定包名
	return nil
}