package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestDree(t *testing.T) {
	var tests []*exec.Cmd
	n := "207598"
	d := "../data/test.db"
	test := exec.Command("./dree", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-n", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-g", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-l", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-l", "-g", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-l", "-n", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-l", "-g", "-n", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-L", "complete", n, d)
	tests = append(tests, test)
	test = exec.Command("./dree", "-L", "complete,chromosome", n, d)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Errorf("couldn't run %q", test)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Errorf("couldn't open %q", f)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s", get, want)
		}
	}
}
