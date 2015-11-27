package sysutil

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGetwd(t *testing.T) {
	dir, error := ioutil.TempDir("", "elvishtest.")
	if error != nil {
		t.Errorf("Got error when creating temp dir: %v", error)
	} else {
		os.Chdir(dir)
		dir, error = filepath.EvalSymlinks(dir)
		if gotwd := Getwd(); gotwd != dir || error != nil {
			t.Errorf("Getwd() -> %v, want %v", gotwd, dir)
		}
		os.Remove(dir)
	}

	home := os.Getenv("HOME")
	os.Chdir(home)
	if gotwd := Getwd(); gotwd != home {
		t.Errorf("Getwd() -> %v, want %v", gotwd, home)
	}
}
