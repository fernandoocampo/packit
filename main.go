package main

import "github.com/fernandoocampo/packit/version"

// PrintHello print hello using the buildin function println().
func PrintHello() {
	println("hello", version.Version)
}

func main() {
	PrintHello()
}
