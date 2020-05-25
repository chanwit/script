package script

import (
	"gopkg.in/pipe.v2"
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
	var name, value string
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

	b := NewBuilder()
	b.pipes = append(b.pipes, pipe.SetEnvVar(name, value))
	return b
}
