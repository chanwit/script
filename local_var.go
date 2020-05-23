package script

import (
	"bytes"
	"strconv"
	"strings"
)

type LocalVar struct {
	buffer *bytes.Buffer
}

func (v *LocalVar) RawString() string {
	return v.buffer.String()
}

func (v *LocalVar) String() string {
	return strings.TrimSpace(v.buffer.String())
}

func (v *LocalVar) Int() int {
	val := v.String()
	i, _ := strconv.Atoi(val)
	return i
}

func (v *LocalVar) Bool() bool {
	val := v.String()
	if val == "1" || val == "true" {
		return true
	}
	return false
}

func (v *LocalVar) Write(p []byte) (n int, err error) {
	return v.buffer.Write(p)
}

func Var() *LocalVar {
	return &LocalVar{
		&bytes.Buffer{},
	}
}
