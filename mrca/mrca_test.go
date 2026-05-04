package main

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestMrca(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.nwk"
	p := "./mrca"
	test := exec.Command(p, "a", f)
	tests = append(tests, test)
	test = exec.Command(p, "e|f", f)
	tests = append(tests, test)
	test = exec.Command(p, "a|c|e", f)
	tests = append(tests, test)
	f = "../data/eco7k.nwk"
	test = exec.Command(p, "941322", f)
	tests = append(tests, test)
	for i, test := range tests {
		g, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		w, err := os.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(g, w) {
			t.Errorf("%s - get:\n%s\nwant:\n%s\n", f, g, w)
		}
	}
}
