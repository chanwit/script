package script

import (
	"fmt"
	"io"
	"os"
	"strings"

	"gopkg.in/pipe.v2"
)

type Builder struct {
	pipes []pipe.Pipe
}

func NewBuilder() *Builder {
	return &Builder{[]pipe.Pipe{}}
}

func (b *Builder) Exec(name string, args ...string) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Build.Exec %s %s\n", name, strings.Join(args, " "))
	}

	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}

func (b *Builder) Tee(w io.Writer) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Build.Tee %v\n", w)
	}

	b.pipes = append(b.pipes, pipe.Tee(w))
	return b
}

func (b *Builder) WriteFile(path string, perm os.FileMode) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Build.WriteFile %s %#o\n", path, perm)
	}

	b.pipes = append(b.pipes, pipe.WriteFile(path, perm))
	return b
}

func (b *Builder) Run() error {
	return pipe.Run(pipe.Line(b.pipes...))
}

func (b *Builder) CombinedOutput() ([]byte, error) {
	return pipe.CombinedOutput(pipe.Line(b.pipes...))
}

func (b *Builder) DividedOutput() ([]byte, []byte, error) {
	return pipe.DividedOutput(pipe.Line(b.pipes...))
}

func (b *Builder) To(w io.Writer) error {
	pipes := append(b.pipes, pipe.Write(w))
	return pipe.Run(pipe.Line(pipes...))
}

func extractExport(i interface{}) (name string, value string) {
	switch v := i.(type) {
	case string:
		parts := strings.SplitN(v, "=", 2)
		if len(parts) == 2 {
			name = parts[0]
			value = parts[1]
		}
	case Exportable:
		name, value = v.Export()
	}
	return
}

func (b *Builder) Export(i interface{}) *Builder {
	if Debug {
		fmt.Printf("[DEBUG] Build.Export %v\n", i)
	}

	name, value := extractExport(i)
	b.pipes = append(b.pipes, pipe.SetEnvVar(name, value))
	return b
}
