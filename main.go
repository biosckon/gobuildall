package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	env := os.Environ()

	archs := csv.NewReader(strings.NewReader(archs_string))

	for {
		rec, err := archs.Read()
		if err == io.EOF {
			break
		}

		if len(rec) != 2 {
			continue
		}

		goos := rec[0]
		goarch := rec[1]

		goos_env := strings.Join([]string{"GOOS", goos}, "=")
		goarch_env := strings.Join([]string{"GOARCH", goarch}, "=")

		new_env := append(env, goos_env, goarch_env)

		wdname, err := os.Getwd()
		if err != nil {
			log.Fatal("Error with os.Getwd() ", err)
		}
		dname, fname := filepath.Split(wdname)

		full_name := filepath.Join(dname, fname, "build", fname)

		log.Println("Building: ", full_name+"_"+goos+"_"+goarch)

		cmd := exec.Command("go", "build", "-o", full_name+"_"+goos+"_"+goarch)
		cmd.Env = new_env

		out, err := cmd.CombinedOutput()
		if err != nil {
			log.Println("Error with starting the command: ", err)
		}
		log.Printf("%s\n", out)

	}

	log.Print("DONE")

}

var archs_string = `
android,arm
darwin,386
darwin,amd64
darwin,arm
darwin,arm64
dragonfly,amd64
freebsd,386
freebsd,amd64
freebsd,arm
linux,386
linux,amd64
linux,arm
linux,arm64
linux,ppc64
linux,ppc64le
linux,mips
linux,mipsle
linux,mips64
linux,mips64le
netbsd,386
netbsd,amd64
netbsd,arm
openbsd,386
openbsd,amd64
openbsd,arm
plan9,386
plan9,amd64
solaris,amd64
windows,386
windows,amd64
`
