package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

const usage = `
  Usage: piccolo [options] <interval> <command>
`
const green1 = "\033[0;40;32m"

const (
	normal        = "\033[0m"
	blod          = "\033[1m"
	italics       = "\033[3m" // test failed
	underline     = "\033[4m"
	inverse       = "\033[7m"
	strikethrough = "\033[9m" // test failed

	foreBlack   = "\033[30m"
	foreRed     = "\033[31m"
	foreGreen   = "\033[32m"
	foreYellow  = "\033[33m"
	foreBlue    = "\033[34m"
	forePurple  = "\033[35m"
	foreCyan    = "\033[36m"
	foreWhite   = "\033[37m"
	foreDefault = "\033[39m"

	backBlack   = "\033[40m"
	backRed     = "\033[41m"
	backGreen   = "\033[42m"
	backYellow  = "\033[43m"
	backBlue    = "\033[44m"
	backPurple  = "\033[45m"
	backCyan    = "\033[46m"
	backWhite   = "\033[47m"
	backDefault = "\033[49m"
)

var (
	flags = flag.NewFlagSet("every", flag.ContinueOnError)
	exit  = flags.Bool("exit", false, "")
)

func printUsage() {
	fmt.Println(usage)
	os.Exit(0)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("%sError: %s", foreRed, err)
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

	log.Printf("%severy %s running %s", forePurple, interval, cmd)

	for {
		time.Sleep(interval)
		start := time.Now()
		log.Printf("%sexec `%s`", foreGreen, cmd)
		proc := exec.Command("/bin/sh", "-c", cmd)

		proc.Start()
		err := proc.Wait()
		ps := proc.ProcessState

		if err != nil {
			log.Printf("pid %d failed with %s", ps.Pid(), ps.String())
			if *exit {
				os.Exit(1)
			}
			continue
		}
		log.Printf("pid %d completed in %s", ps.Pid(), time.Since(start))
	}
}
