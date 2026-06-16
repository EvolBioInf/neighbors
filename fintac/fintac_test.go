package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestFintac(t *testing.T) {
	var tests []*exec.Cmd
	f := "./test.nwk"
	p := "./fintac"
	test := exec.Command(p, f)
	tests = append(tests, test)
	test = exec.Command(p, "-a", f)
	tests = append(tests, test)
	test = exec.Command(p, "-t", "^tGC[AF]", f)
	tests = append(tests, test)
	test = exec.Command(p, "-n", "^nGC[AF]", f)
	tests = append(tests, test)
	test = exec.Command(p, "-u", "^n", f)
	tests = append(tests, test)
	test = exec.Command(p, "-w", "0", f)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		name := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := os.ReadFile(name)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
