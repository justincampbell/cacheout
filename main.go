package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/justincampbell/cacheout/cache"
)

const usage = `
usage: cacheout [duration] [command]
example durations: 5s 15m 1h30m
`

// A Command holds the parsed input argument from the user
type Command struct {
	key  string
	ttl  time.Duration
	bin  string
	args []string
}

func init() {
	flag.Parse()
}

func main() {
	command, err := parseArgs(flag.Args())
	if err != nil {
		fmt.Printf("%s%s", err, usage)
		os.Exit(1)
	}

	fc := cache.NewFileCache(command.key, &command.ttl)

	if fc.Stale() {
		err := command.execAndCache(fc)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		_, err = os.Stdout.Write(fc.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func parseArgs(args []string) (Command, error) {
	command := Command{}

	if len(args) < 2 {
		return command, fmt.Errorf("not enough arguments")
	}

	command.key = hashCommand(args)

	ttlArg, bin, args := args[0], args[1], args[2:]
	command.bin = bin
	command.args = args

	ttl, err := time.ParseDuration(ttlArg)
	if err != nil {
		return command, err
	}
	command.ttl = ttl

	return command, nil
}

func hashCommand(command []string) string {
	h := md5.New()

	for _, part := range command {
		_, err := io.WriteString(h, part)
		if err != nil {
			log.Fatal(err)
		}
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

func (command *Command) execAndCache(cache cache.Cache) error {
	cmd := exec.Command(command.bin, command.args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = io.Copy(os.Stdout, io.TeeReader(stdout, cache))
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if exiterr, ok := err.(*exec.ExitError); ok {
		if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
			os.Exit(status.ExitStatus())
		} else {
			return err
		}
	}

	return cache.Persist()
}
