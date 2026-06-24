package clusters

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"testing"
)

func TestClusters(t *testing.T) {
	var tests []*exec.Cmd
	f := "test.nwk"
	p := "./cmd/clusters"
	test := exec.Command(p, f)
	tests = append(tests, test)
	test = exec.Command(p, "-b", "3090", f)
	tests = append(tests, test)
	test = exec.Command(p, "-c", f)
	tests = append(tests, test)
	test = exec.Command(p, "-f", "1.5", f)
	tests = append(tests, test)
	test = exec.Command(p, "-s", "4", f)
	tests = append(tests, test)
	test = exec.Command(p, "-t", f)
	tests = append(tests, test)
	test = exec.Command(p, "-T", f)
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
			t.Errorf("%s - get:\n%s\nwant:\n%s\n", f, get, want)
		}
	}
}
