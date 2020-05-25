package script

import (
	"fmt"
)

func Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Printf(format, a...)
}

func Print(a ...interface{}) (n int, err error) {
	return fmt.Print(a...)
}

func Println(a ...interface{}) (n int, err error) {
	return fmt.Println(a...)
}
