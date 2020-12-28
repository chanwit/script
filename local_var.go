package script

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"
)

type LocalVar struct {
	name   string
	buffer *bytes.Buffer
}

type Exportable interface {
	Export() (string, string)
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

func (v *LocalVar) Export() (string, string) {
	return v.name, v.String()
}

func (v *LocalVar) Lines() (lines []string) {
	scanner := bufio.NewScanner(v.buffer)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Var() *LocalVar {
	return &LocalVar{
		"",
		&bytes.Buffer{},
	}
}

func NamedVar(name string) *LocalVar {
	return &LocalVar{
		name,
		&bytes.Buffer{},
	}
}
