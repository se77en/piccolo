package main

import (
	"flag"
	"fmt"
	"github.com/se77en/piccolo/piccolo"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const usage = `
  Usage: piccolo <interval> <command>
`

var (
	flags = flag.NewFlagSet("piccolo", flag.ContinueOnError)
	exit  = flags.Bool("exit", false, "")
)

type timingFunc struct {
	Func   func()
	Ticker int
}

func printUsage() {
	fmt.Println(usage)
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(piccolo.SetTextColor("red", fmt.Sprintf("Error: %s", err)))
	}
}

func execCmd(cmd string) func() {
	return func() {
		start := time.Now()
		log.Printf("exec %s", cmd)
		proc := exec.Command("/bin/sh", "-c", cmd)

		proc.Start()
		err := proc.Wait()
		ps := proc.ProcessState

		if err != nil {
			log.Printf("pid %d failed with %s", ps.Pid(), ps.String())
			if *exit {
				os.Exit(1)
			}
		}
		log.Printf("pid %d completed in %s", ps.Pid(), time.Since(start))
	}
}

func main() {
	flags.Usage = printUsage
	flags.Parse(os.Args[1:])
	argv := flags.Args()

	if len(argv) < 1 {
		checkErr(fmt.Errorf("<interval> required"))
	}

	interval, err := time.ParseDuration(argv[0])
	checkErr(err)

	if len(argv) < 2 {
		checkErr(fmt.Errorf("<command> required"))
	}

	cmd := strings.Join(argv[1:], " ")

	log.Print(piccolo.SetTextColor("magenta", fmt.Sprintf("every %s running %s", interval, cmd)))

	piccolo.AddTimingFunc("Command Line Job", -1, execCmd(cmd))
	piccolo.StartTiming(interval)
}
