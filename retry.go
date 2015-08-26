package main

import (
	"flag"
	"fmt"
	"github.com/cenkalti/backoff"
	"log"
	"os"
	"os/exec"
	"time"
)

var (
	flagInitialInterval = flag.Int("initialInterval", 1, "retry interval(s)")
	flagMaxElapsedTime  = flag.Int("maxElapsedTime", 10000, "Max Elapsed Time(s) is limit of backoff steps. If the job spends over this, job makes stopped. If set 0, the job will never stop.")
	flagMaxInterval     = flag.Int("maxInterval", 1000, "cap of retry interval(s)")
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: retry <command>\n\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		usage()
	}

	var b []byte
	operation := func() error {
		var err error
		b, err = exec.Command(flag.Arg(0), args[1:]...).Output()
		if err != nil {
			log.Printf("err: %s", err)
		}
		return err
	}

	bf := backoff.NewExponentialBackOff()
	second := func(i int) time.Duration {
		return time.Duration(i) * time.Second
	}

	bf.MaxElapsedTime = second(*flagMaxElapsedTime)
	bf.MaxInterval = second(*flagMaxInterval)
	bf.InitialInterval = second(*flagInitialInterval)

	err := backoff.Retry(operation, bf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "operation failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, string(b))
	os.Exit(0)
}
