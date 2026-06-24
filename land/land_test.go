package land

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestLand(t *testing.T) {
	var tests []*exec.Cmd
	f := "test1.nwk"
	p := "./cmd/land"
	test := exec.Command(p, f)
	tests = append(tests, test)
	test = exec.Command(p, "-p", "p", f)
	tests = append(tests, test)
	test = exec.Command(p, "-s", "s", f)
	tests = append(tests, test)
	test = exec.Command(p, "-l", f)
	tests = append(tests, test)
	f = "test2.nwk"
	test = exec.Command(p, "-r", f)
	tests = append(tests, test)
	for i, test := range tests {
		get, err := test.Output()
		if err != nil {
			t.Error(err)
		}
		f := "r" + strconv.Itoa(i+1) + ".txt"
		want, err := ioutil.ReadFile(f)
		if err != nil {
			t.Error(err)
		}
		if !bytes.Equal(get, want) {
			t.Errorf("get:\n%s\nwant:\n%s\n", get, want)
		}
	}
}
