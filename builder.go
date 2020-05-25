package script

import (
	"io"
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
	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}

func (b *Builder) Tee(w io.Writer) *Builder {
	b.pipes = append(b.pipes, pipe.Tee(w))
	return b
}

func (b *Builder) Run() error {
	return pipe.Run(pipe.Line(b.pipes...))
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
	name, value := extractExport(i)
	b.pipes = append(b.pipes, pipe.SetEnvVar(name, value))
	return b
}
