package pickle

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strconv"
	"testing"
)

func TestPickle(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.nwk"
	p := "./cmd/pickle"
	test := exec.Command(p, "7", f)
	tests = append(tests, test)
	test = exec.Command(p, "7,3", f)
	tests = append(tests, test)
	test = exec.Command(p, "9", f)
	tests = append(tests, test)
	test = exec.Command(p, "4", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "7", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "7,3", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "9", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "4", f)
	tests = append(tests, test)
	test = exec.Command(p, "-t", "7", f)
	tests = append(tests, test)
	test = exec.Command(p, "-t", "7,3", f)
	tests = append(tests, test)
	test = exec.Command(p, "-t", "9", f)
	tests = append(tests, test)
	test = exec.Command(p, "-t", "4", f)
	tests = append(tests, test)
	test = exec.Command(p, "-C", "7", f)
	tests = append(tests, test)
	test = exec.Command(p, "-C", "7,3", f)
	tests = append(tests, test)
	test = exec.Command(p, "-C", "9", f)
	tests = append(tests, test)
	test = exec.Command(p, "-C", "4", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "-t", "7", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "-t", "7,3", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "-t", "9", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", "-t", "4", f)
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
