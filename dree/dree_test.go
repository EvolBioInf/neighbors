package dree

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

type Test struct {
	t *exec.Cmd
	r int
}

func TestDree(t *testing.T) {
	var tests []Test
	n := "207598"
	d := "../data/test.db"
	p := "./cmd/dree"
	test := Test{t: exec.Command(p, n, d), r: 1}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		n, d), r: 1}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-n", n, d), r: 2}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-n", n, d), r: 2}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-g", n, d), r: 3}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-g", n, d), r: 3}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-l", n, d), r: 4}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-l", n, d), r: 4}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-l", "-g", n, d), r: 5}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-l", "-g", n, d), r: 5}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-l", "-n", n, d), r: 6}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-l", "-n", n, d), r: 6}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-l", "-g", "-n", n, d), r: 7}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-l", "-g", "-n", n, d), r: 7}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-L", "complete", n, d), r: 8}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-L", "complete", n, d), r: 8}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-L", "complete,chromosome", n, d), r: 9}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-L", "complete,chromosome", n, d), r: 9}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-m", "-2", n, d), r: 10}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-m", "-2", n, d), r: 10}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-m", "0", n, d), r: 11}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-m", "0", n, d), r: 11}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-m", "2", n, d), r: 12}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-m", "2", n, d), r: 12}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-m", "200", n, d), r: 13}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-m", "200", n, d), r: 13}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-R", "-l", n, d), r: 14}
	tests = append(tests, test)
	test = Test{t: exec.Command(p, "-D", "test.db",
		"-R", "-l", n, d), r: 14}
	tests = append(tests, test)
	for _, test := range tests {
		get, err := test.t.Output()
		if err != nil {
			t.Errorf("couldn't run %q", test.t)
		}
		f := "r" + strconv.Itoa(test.r) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Errorf("couldn't open %q", f)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s", get, want)
		}
	}
}
