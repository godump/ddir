package ddir

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

var (
	cBase string
	cDone bool
)

func Base(base string) {
	if cDone {
		log.Fatalln("ddir: call Base() multiple times")
	}
	cBase = base
	cDone = true
	Make()
}

func Auto(elem ...string) {
	seps := []string{}
	if runtime.GOOS == "windows" {
		seps = append(seps, os.Getenv("localappdata"))
		seps = append(seps, elem...)
		Base(filepath.Join(seps...))
		return
	}
	if runtime.GOOS == "linux" && runtime.GOARCH == "arm" {
		return
	}
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	seps = append(seps, u.HomeDir)
	seps = append(seps, elem...)
	seps[1] = "." + seps[1]
	Base(filepath.Join(seps...))
}

func Join(elem ...string) string {
	return filepath.Join(append([]string{cBase}, elem...)...)
}

func Make(elem ...string) {
	err := os.Mkdir(Join(elem...), 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalln(err)
	}
}
