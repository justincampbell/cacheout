package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
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

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}

		_, err = os.Stdout.Write(output)
		if err != nil {
			log.Fatal(err)
		}

		fc.Write(output)
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
