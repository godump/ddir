package ddir

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	base string
	done bool
)

// Base sets base data path. This should be called from a binary's init or
// main function. If the base is already set then this will log.Panicln.
func Base(base string) {
	if done {
		log.Panicln("ddir: call Base() multiple times")
	}
	base = base
	done = true
	Make()
}

// Join joins any number of path elements into a single path, adding
// a Separator if necessary.
func Join(elem ...string) string {
	return filepath.Join(append([]string{base}, elem...)...)
}

func none(name string) bool {
	_, err := os.Stat(name)
	return err != nil && os.IsNotExist(err)
}

// Make ensures the directory structure. If folder doesn’t exist, create it.
//
// Example:
//   Base("/tmp/")
//   Make("data0") // make /tmp/data0 if it doesn't exist
//   Make("data1") // make /tmp/data1 if it doesn't exist
func Make(elem ...string) {
	name := Join(elem...)
	if !none(name) {
		return
	}
	if err := os.Mkdir(name, 0755); err != nil {
		log.Panicln(err)
	}
}

func autoMakeWindows(elem ...string) {
	seps := []string{}
	seps = append(seps, os.Getenv("localappdata"))
	seps = append(seps, elem...)
	Base(filepath.Join(seps...))
}

func autoMakeAndroid(elem ...string) {
}

func autoMakeDefault(elem ...string) {
	u, err := user.Current()
	if err != nil {
		log.Panicln(err)
	}
	seps := []string{}
	seps = append(seps, u.HomeDir)
	seps = append(seps, elem...)
	seps[1] = "." + strings.ToLower(seps[1])
	Base(filepath.Join(seps...))
}

// Auto is an automatic Base function call affected by the operating system.
// Most applications' data directories follow the operating system's
// specifications, for example, the data directory of vim is placed in ~/.vim.
//
// Example:
//   Auto("play") // Equals with Base("~/AppData/Local/play") on windows
//   Auto("play") // Equals with Base("~/.play") on linux
func Auto(elem ...string) {
	switch {
	case runtime.GOOS == "windows":
		autoMakeWindows(elem...)
	case runtime.GOOS == "linux" && runtime.GOARCH == "arm":
		autoMakeAndroid(elem...)
	default:
		autoMakeDefault(elem...)
	}
}
