package main

import (
	"fmt"
	"strings"

	"github.com/fernandoocampo/packit/version"
	"golang.org/x/sys/unix"
)

// PrintHello print hello using the buildin function println().
func PrintHello() {
	println("hello", version.Version)
}

func main() {
	PrintHello()
	uname := unix.Utsname{}
	unix.Uname(&uname)
	fmt.Println(arrayToString(uname.Nodename[:]))
	fmt.Println(arrayToString(uname.Release[:]))
	fmt.Println(arrayToString(uname.Sysname[:]))
	fmt.Println(arrayToString(uname.Version[:]))
	fmt.Println(arrayToString(uname.Machine[:]))
}

func arrayToString(x []byte) string {
	buf := make([]byte, len(x))
	for i, b := range x {
		buf[i] = byte(b)
	}
	str := string(buf[:])
	if i := strings.Index(str, "\x00"); i != -1 {
		str = str[:i]
	}
	return str
}
