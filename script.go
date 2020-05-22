package script

import (
	"bytes"
	"strconv"
	"strings"

	"gopkg.in/pipe.v2"
)

type Var struct {
	buffer *bytes.Buffer
}

func (v *Var) RawString() string {
	return v.buffer.String()
}

func (v *Var) String() string {
	return strings.TrimSpace(v.buffer.String())
}

func (v *Var) Int() int {
	val := v.String()
	i, _ := strconv.Atoi(val)
	return i
}

func (v *Var) Bool() bool {
	val := v.String()
	if val == "1" || val == "true" {
		return true
	}
	return false
}

func NewVar() *Var {
	return &Var{
		&bytes.Buffer{},
	}
}

type Builder struct {
	pipes []pipe.Pipe
}

func (b *Builder) Exec(name string, args ...string) *Builder {
	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}

func (b *Builder) Tee(v *Var) *Builder {
	b.pipes = append(b.pipes, pipe.Tee(v.buffer))
	return b
}

func (b *Builder) Run() error {
	return pipe.Run(pipe.Line(b.pipes...))
}

func (b *Builder) To(v *Var) error {
	output, err := pipe.Output(pipe.Line(b.pipes...))
	v.buffer = bytes.NewBuffer(output)
	return err
}

func NewBuilder() *Builder {
	return &Builder{[]pipe.Pipe{}}
}
