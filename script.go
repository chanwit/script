package script

import (
	"gopkg.in/pipe.v2"
	"os"
	"os/exec"
	"strings"
)

func Exec(name string, args ...string) *Builder {
	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}

func Echo(args ...interface{}) *Builder {
	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.Println(args...))
	return b
}

func Export(i interface{}) *Builder {
	name, value := extractExport(i)
	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.SetEnvVar(name, value))
	return b
}

func System(cmd string) *Builder {
	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.System(cmd))
	return b
}

func Sudo(args ...string) error {
	sudoWithArgs := strings.Join(append([]string{"sudo"}, args...), " ")
	cmd := exec.Command("sh", "-c", sudoWithArgs)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}