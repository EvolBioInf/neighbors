package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestClimt(t *testing.T) {
	tests := make([]*exec.Cmd, 0)
	f := "test.nwk"
	s := "303"
	test := exec.Command("./climt", s, f)
	tests = append(tests, test)
	test = exec.Command("./climt", "-d", s, f)
	tests = append(tests, test)
	s = "^30[34]$"
	test = exec.Command("./climt", "-r", s, f)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
