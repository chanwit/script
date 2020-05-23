package script

import (
	"os"
)

func Stdout() *os.File {
	return os.Stdout
}

func Stdin() *os.File {
	return os.Stdin
}

func Stderr() *os.File {
	return os.Stderr
}