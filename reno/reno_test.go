package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

func TestReno(t *testing.T) {
	test := exec.Command("./reno", "names.txt", "reno1.nwk")
	get, err := test.Output()
	if err != nil {
		t.Error(err)
	}
	want, err := os.ReadFile("reno2.nwk")
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(get, want) {
		t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
	}
}
