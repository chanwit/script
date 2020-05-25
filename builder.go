package script

import (
	"io"

	"gopkg.in/pipe.v2"
)

type Builder struct {
	pipes []pipe.Pipe
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

func NewBuilder() *Builder {
	return &Builder{[]pipe.Pipe{}}
}
