package script

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/pipe.v2"
)

var Debug bool = false

func Exec(name string, args ...string) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Exec %s %s\n", name, strings.Join(args, " "))
	}

	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}

func Echo(args ...interface{}) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Echo %v\n", args)
	}

	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.Println(args...))
	return b
}

func Export(i interface{}) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Export %v\n", i)
	}

	name, value := extractExport(i)
	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.SetEnvVar(name, value))
	return b
}

func System(cmd string) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] System %s\n", cmd)
	}

	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.System(cmd))
	return b
}

func Sudo(args ...string) error {
	sudoWithArgs := strings.Join(append([]string{"sudo"}, args...), " ")
	if Debug {
		fmt.Printf("[DEBUG] Sudo sh -c %s\n", sudoWithArgs)
	}

	cmd := exec.Command("sh", "-c", sudoWithArgs)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Run(name string, args ...string) error {
	if Debug {
		fmt.Printf("[DEBUG] Run %s %s\n", name, strings.Join(args, " "))
	}

	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
