package easyfs

import (
	"os"
	"path/filepath"
)

var (
	base string
	done bool
)

// Base sets base data path. This should be called from a binary's init or main function.
func Base(name string) {
	if done {
		panic("res: call Base() multiple times")
	}
	if name == "" {
		base, _ = filepath.Abs(".")
	} else {
		base, _ = filepath.Abs(name)
	}
	done = true
	Make("")
}

// Base on path name for the executable that started the current process.
func BaseExec() {
	d, err := os.Executable()
	if err != nil {
		panic(err)
	}
	Base(filepath.Dir(d))
}

// Path joins any number of path elements into a single path, adding a Separator if necessary.
func Path(name string) string {
	return filepath.Join(base, name)
}

// Make ensures the directory structure. If folder doesn’t exist, create it.
//
// Example:
//   Base("/tmp")
//   Make("/data0") // make /tmp/data0 if it doesn't exist
//   Make("/data1") // make /tmp/data1 if it doesn't exist
func Make(name string) {
	if _, err := os.Stat(Path(name)); err == nil || os.IsExist(err) {
		return
	}
	if err := os.MkdirAll(Path(name), 0755); err != nil {
		panic(err)
	}
}
