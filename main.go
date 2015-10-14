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

func init() {
	flag.Parse()
}

func main() {
	command := flag.Args()
	key := hashCommand(command)
	ttlArg, bin, args := command[0], command[1], command[2:]

	ttl, err := time.ParseDuration(ttlArg)
	if err != nil {
		log.Fatal(err)
	}

	fc := cache.NewFileCache(key, &ttl)

	if fc.Stale() {
		cmd := exec.Command(bin, args...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			log.Fatal(err)
		}

		if err := cmd.Start(); err != nil {
			log.Fatal(err)
		}

		_, err = io.Copy(os.Stdout, io.TeeReader(stdout, fc))
		if err != nil {
			log.Fatal(err)
		}

		err = cmd.Wait()
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			} else {
				log.Fatal(err)
			}
		}

		fc.Persist()
	} else {
		_, err = os.Stdout.Write(fc.Bytes())
		if err != nil {
			log.Fatal(err)
		}
	}
}

func hashCommand(command []string) string {
	h := md5.New()

	for _, part := range command {
		io.WriteString(h, part)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}
