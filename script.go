package script

import (
	"gopkg.in/pipe.v2"
)

func Exec(name string, args ...string) *Builder {
	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.Exec(name, args...))
	return b
}
