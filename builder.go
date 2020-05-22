package script

import (
	"bytes"
	"gopkg.in/pipe.v2"
)

type Builder struct {
	pipes []pipe.Pipe
}

func (b *Builder) Exec(name string, args ...string) *Builder {
	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}

func (b *Builder) Tee(v *LocalVar) *Builder {
	b.pipes = append(b.pipes, pipe.Tee(v.buffer))
	return b
}

func (b *Builder) Run() error {
	return pipe.Run(pipe.Line(b.pipes...))
}

func (b *Builder) To(v *LocalVar) error {
	output, err := pipe.Output(pipe.Line(b.pipes...))
	v.buffer = bytes.NewBuffer(output)
	return err
}

func NewBuilder() *Builder {
	return &Builder{[]pipe.Pipe{}}
}

