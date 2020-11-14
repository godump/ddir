package easyfs

import (
	"testing"
)

func TestEasyFs(t *testing.T) {
	Base(".")
	t.Log(Path("/a"))
}
