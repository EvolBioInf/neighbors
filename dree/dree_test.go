package dree

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
	p := "./cmd/dree"
	test := exec.Command(p, n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-n", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-g", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-l", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-l", "-g", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-l", "-n", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-l", "-g", "-n", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-L", "complete", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-L", "complete,chromosome", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-m", "-2", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-m", "0", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-m", "2", n, d)
	tests = append(tests, test)
	test = exec.Command(p, "-m", "200", n, d)
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
