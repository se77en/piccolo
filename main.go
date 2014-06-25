package main

import (
	"fmt"
)

const usage = `
  Usage: piccolo [options] <command> <interval>
`

func main() {
	fmt.Println(usage)
	fmt.Println("Hello Piccolo")
}
